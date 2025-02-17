package account

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nedpals/supabase-go"
	"log/slog"
	"s-coder-snippet-sharder/api/apierror"
	"s-coder-snippet-sharder/api/cookieutil"
	"s-coder-snippet-sharder/api/handler"
	"s-coder-snippet-sharder/internal/app"
	"s-coder-snippet-sharder/internal/app/account"
	"s-coder-snippet-sharder/pkg/goth_fiber"
	sessionutil "s-coder-snippet-sharder/pkg/session"
	"time"
)

// Regieter user router
func Routes(h *handler.Handler) {
	authGroup := h.App.Group("/api")
	{
		authGroup.Post("/signup", SignUpUser(h))
		authGroup.Get("/auth/logout", LogoutUser(h))
		authGroup.Get("/auth/:provider", Login(h))
		authGroup.Get("/auth/:provider/callback", HandleauthCallBackFunction(h))
	}
}

func Login(h *handler.Handler) fiber.Handler {
	return goth_fiber.BeginAuthHandler
}

func HandleauthCallBackFunction(h *handler.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, err := goth_fiber.CompleteUserAuth(ctx)
		ctxt := ctx.Context()
		if err != nil {
			slog.ErrorContext(ctxt, "ComplateteUserAuth error: %v", err)
			return apierror.ErrorResponse(ctx, apierror.ErrorInternal)
		}

		duration, _ := time.ParseDuration("1h")
		token, _, err := h.Token.CreateToken(user.UserID, user.Name, duration)
		if err != nil {
			slog.ErrorContext(ctxt, "CreateToken error", "error", err)
			return apierror.ErrorResponse(ctx, apierror.ErrorInternal)
		}

		if err := sessionutil.Set(ctx, user.UserID, user); err != nil {
			slog.ErrorContext(ctxt, "Set session error", "error", err)
			return apierror.ErrorResponse(ctx, apierror.ErrorInternal)
		}

		cookieutil.SetAccessToken(ctx, token, 24*60*60)

		// Redirect về trang chủ
		return ctx.Redirect("/", fiber.StatusSeeOther)
	}
}

func LogoutUser(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := goth_fiber.Logout(c); err != nil {
			slog.ErrorContext(c.Context(), "Logout error", "error", err)
			return apierror.ErrorResponse(c, apierror.ErrorInternal)
		}

		cookieutil.SetAccessToken(c, "", -1)
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}
}

func SignUpUser(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()
		var data account.User

		if err := c.BodyParser(&data); err != nil {
			slog.ErrorContext(ctx, "failed to parse data", "error", err)
			return apierror.ErrorResponse(c, apierror.ErrorInternal)
		}

		user, err := h.AccountAPI.SignUp(ctx, data.Email, data.Password)
		if err != nil {
			if errors.Is(err, app.ErrInvalidInput) {
				return apierror.ErrorResponse(c, err)
			}
			return apierror.ErrorResponse(c, apierror.ErrorInternal)
		}

		_, err = h.SB.Auth.SignUp(ctx, supabase.UserCredentials{
			Email:    user.Email,
			Password: user.Password,
		})
		if err != nil {
			slog.ErrorContext(ctx, "failed to sign up with supabase", "error", err)
			return apierror.ErrorResponse(c, err)
		}

		details, err := h.SB.Auth.SignIn(ctx, supabase.UserCredentials{
			Email:    user.Email,
			Password: user.Password,
		})
		if err != nil {
			slog.ErrorContext(ctx, "failed to sign in", "error", err)
			return apierror.ErrorResponse(c, errors.New("Something went wrong, failed to sign in"))
		}
		duration, _ := time.ParseDuration("1h")
		token, _, err := h.Token.CreateToken(details.User.ID, details.User.Role, duration)
		// Tạo cookie chứa JWT
		cookie := fiber.Cookie{
			Name:     "cookieName",
			Value:    token,
			Expires:  time.Now().Add(24 * time.Hour), // Cookie hết hạn sau 24 giờ
			HTTPOnly: true,                           // Ngăn chặn truy cập từ JavaScript
			Secure:   true,                           // Chỉ gửi cookie qua HTTPS
			Path:     "/",                            // Áp dụng cho toàn bộ site
		}

		c.Cookie(&cookie)

		//return templateutil.Render(h.View, c, "snippet/main", fiber.Map{
		//	"User": details.User,
		//})
		return c.Status(fiber.StatusOK).JSON(details.User)
	}
}
