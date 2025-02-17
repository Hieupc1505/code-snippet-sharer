package cookieutil

import (
	"github.com/gofiber/fiber/v2"
	"s-coder-snippet-sharder/types"
)

func SetRefreshToken(c *fiber.Ctx, refreshToken string) {
	c.Cookie(&fiber.Cookie{
		Name:     types.RefreshTokenKey,
		SameSite: "lax",
		Value:    refreshToken,
		Path:     "/",
		HTTPOnly: true,
	})
}

// Set asscess token
func SetAccessToken(c *fiber.Ctx, accessToken string, expiresIn int) {
	c.Cookie(&fiber.Cookie{
		Name:     types.CookieToken,
		SameSite: "lax",
		MaxAge:   expiresIn,
		Value:    accessToken,
		Path:     "/",
		HTTPOnly: true,
	})
}
