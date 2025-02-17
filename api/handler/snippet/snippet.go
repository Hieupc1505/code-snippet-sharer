package snippet

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"s-coder-snippet-sharder/api/apierror"
	"s-coder-snippet-sharder/api/handler"
	mid "s-coder-snippet-sharder/api/middleware"
	"s-coder-snippet-sharder/api/templateutil"
	"s-coder-snippet-sharder/internal/app"
	db "s-coder-snippet-sharder/internal/db/sqlc"
	"s-coder-snippet-sharder/types"
)

func Routes(h *handler.Handler) {
	h.App.Get("/", mid.WithAuthenticatedUser(h), GetMain(h))

	snippetGroup := h.App.Group("/snippets")
	{
		snippetGroup.Get("/p/:slug", mid.WithAuthenticatedUser(h), ViewPost(h))
		snippetGroup.Get("/recent", GetRecent(h))
	}

	snippetAPIGroup := h.App.Group("/api/snippets")
	{
		snippetAPIGroup.Post("/add", mid.RequireAuthentication(h), AddSnippet(h))
	}
}

func GetRecent(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		snippets, err := h.SnippetAPI.GetPublicSnippets(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get recent snippets", "error", err)
			return apierror.ErrorResponse(c, apierror.ErrorInternal)
		}
		return templateutil.Render(h.View, c, "snippet/recent_post", fiber.Map{
			"Snippets": snippets,
		}, "layouts/none")
	}
}

func AddSnippet(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var data db.AddParams
		if err := c.BodyParser(&data); err != nil {
			return apierror.ErrorResponse(c, err)
		}
		ctx := c.Context()
		snippet, err := h.SnippetAPI.CreateNewCodePost(ctx, data.Lang, data.Title, data.Snippet)
		switch {
		case errors.Is(err, app.ErrInvalidInput):
			slog.ErrorContext(ctx, "failed to create snippet", "error", err)
			return apierror.ErrorResponse(c, err)
		case err != nil:
			slog.ErrorContext(ctx, "failed to create snippet", "error", err)
			return apierror.ErrorResponse(c, apierror.ErrorInternal)
		}
		return templateutil.FlasBackWithSuccess(c, fiber.Map{
			"Snippet": snippet,
		})
	}
}

func GetMain(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals(types.LocalsUserKey)
		return templateutil.Render(h.View, c, "snippet/main", fiber.Map{
			"User": user,
		})
	}
}

func ViewPost(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		slug := c.Params("slug")
		snippet, err := h.SnippetAPI.GetSnippetBySlug(ctx, slug)

		user := c.Locals(types.LocalsUserKey)

		switch {
		case errors.Is(err, app.ErrNotFound):
			slog.ErrorContext(ctx, "failed to get snippet", "error", err)
			// Đặt header Hx-Trigger đúng JSON format
			return apierror.ErrorResponse(c, err)
		case err != nil:
			slog.ErrorContext(ctx, "failed to get snippet", "error", err)
			return apierror.ErrorResponse(c, apierror.ErrorInternal)
		}

		requestorAddr := c.Context().RemoteAddr()
		h.Tasks.Add(1)
		go func(callerIp string) {
			defer h.Tasks.Done()
			//Create a new context, since the request context could cancel before
			//thís operation finishes
			ctx := context.Background()
			if h.IPCache.Has(callerIp) {
				slog.InfoContext(ctx, "IP has been found in the IP cache, not incrementing view count")
			} else {
				// update the view count in the background
				slog.InfoContext(ctx, "About to update view count for slug", "slug", slug)
				if err := h.SnippetRepo.UpdateViewCount(ctx, slug); err != nil {
					slog.ErrorContext(ctx, "failed to update view count", "error", err)
					return
				}
				h.IPCache.Set(callerIp)
			}
		}(requestorAddr.String())

		return templateutil.Render(h.View, c, "snippet/view_post", fiber.Map{
			"Snippet": snippet,
			"User":    user,
		})
	}
}
