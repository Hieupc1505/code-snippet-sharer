package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nedpals/supabase-go"
	"html/template"
	"log/slog"
	"net/http"
	"s-coder-snippet-sharder/api/apierror"
	"s-coder-snippet-sharder/api/templateutil"
	"s-coder-snippet-sharder/internal/app"
	"s-coder-snippet-sharder/internal/app/account"
	"s-coder-snippet-sharder/internal/app/snippet"
	"s-coder-snippet-sharder/pkg/errsx"
	"s-coder-snippet-sharder/pkg/ipcache"
	"s-coder-snippet-sharder/pkg/token"
	"s-coder-snippet-sharder/types"
	"sync"
)

type Handler struct {
	App            *fiber.App
	View           *templateutil.View
	NamedEndpoints templateutil.NamedEndpoints
	Template       *template.Template
	SB             *supabase.Client
	Token          token.Maker
	IPCache        *ipcache.Cache
	Tasks          *sync.WaitGroup

	SnippetRepo snippet.ReadWriter
	SnippetAPI  *snippet.Service

	AccountRepo account.ReadWriter
	AccountAPI  *account.Service
}

func (h *Handler) Logger(c *fiber.Ctx) *slog.Logger {
	value := c.Locals(types.LoggerCtx)
	if value == nil {
		return slog.Default()
	}
	logger, ok := value.(*slog.Logger)
	if !ok {
		panic(fmt.Sprintf("could not assert logger as %T", logger))
	}
	return logger
}

func (h *Handler) RenderComponent(c *fiber.Ctx, template string, err error, data fiber.Map) error {
	logger := h.Logger(c)

	if err != nil {
		if errors.Is(err, app.ErrInvalidInput) {
			var errMap errsx.Map
			if errors.As(err, &errMap) {
				data["ErrorDetails"] = errMap
			}
			data["Error"] = errorHandler(err)
		}
	}
	bytesW := new(bytes.Buffer)
	if err := h.Template.ExecuteTemplate(bytesW, template, data); err != nil {
		logger.ErrorContext(c.Context(), "failed to exec template", "error", err)
		return apierror.ErrorResponse(c, apierror.ErrorInternal)
	}

	c.Set("content-type", "text/html")
	return c.Status(http.StatusOK).SendString(bytesW.String())
}

func errorHandler(err error) string {
	switch {
	case errors.Is(err, app.ErrUnauthorised):
		return "You must login to view this resource"
	case errors.Is(err, app.ErrForbidden):
		return "You do not have permision to access this resource"
	default:
		return "An error has occurred"
	}
}
