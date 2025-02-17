package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"log/slog"
	"s-coder-snippet-sharder/api/apierror"
	"s-coder-snippet-sharder/api/contextutil"
	"s-coder-snippet-sharder/api/handler"
	"s-coder-snippet-sharder/types"
)

// required login to continue
func RequireAuthentication(h *handler.Handler) fiber.Handler {
	slog.Info("RequireAuthentication")
	return func(c *fiber.Ctx) error {

		// Lấy token từ cookie
		token := c.Cookies(types.CookieToken)
		if token == "" {
			return apierror.ErrorResponse(c, apierror.ErrUnauthorised)
		}

		// Xác thực token
		payload, err := h.Token.VerifyToken(token)
		if err != nil {
			return apierror.ErrorResponse(c, apierror.ErrUnauthorised)
		}

		user, err := contextutil.GetUser(c, payload.UserID)
		if err != nil {
			return apierror.ErrorResponse(c, apierror.ErrUnauthorised)
		}

		// Lưu userID vào context
		c.Locals(types.LocalsUserKey, user)

		// Tiếp tục xử lý request
		return c.Next()
	}

}

// add middleware to add user to locals
func WithAuthenticatedUser(h *handler.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals(types.LocalsUserKey, nil)

		//authentication
		token := c.Cookies(types.CookieToken)
		if token == "" {
			return c.Next()
		}

		payload, err := h.Token.VerifyToken(token)
		if err != nil {
			log.Println("Failed to create client: ", err)
			c.ClearCookie(types.CookieToken)
			return c.Next()
		}

		user, err := contextutil.GetUser(c, payload.UserID)
		if err != nil {
			log.Println("Failed to find user: ", err)
			c.ClearCookie(types.CookieToken)
			return c.Next()
		}

		c.Locals(types.LocalsUserKey, user)
		return c.Next()
	}
}
