package contextutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"log"
	sessionutil "s-coder-snippet-sharder/pkg/session"
	"s-coder-snippet-sharder/types"
)

// SetUser lưu user vào context
func SetUser(c *fiber.Ctx, user interface{}) {
	c.Locals(types.LocalsUserKey, user)
}

// GetUser lấy user từ context
func GetUser(c *fiber.Ctx, ID string) (*goth.User, error) {
	user, err := sessionutil.Get(c, ID)
	if err != nil {
		log.Println("Failed to get user from session: ", err)
		c.ClearCookie(types.CookieToken)
		return nil, err
	}

	gothUser, ok := user.(goth.User)
	if !ok {
		log.Println("User is not of type goth.User")
		c.ClearCookie(types.CookieToken)
		return nil, err
	}
	return &gothUser, nil
}

//// SetLogger lưu logger vào context
//func SetLogger(c *fiber.Ctx, logger interface{}) {
//	c.Locals(LoggerCtx, logger)
//}
//
//// GetLogger lấy logger từ context
//func GetLogger(c *fiber.Ctx) interface{} {
//	return c.Locals(LoggerCtx)
//}
//
//// SetPassport lưu passport vào context
//func SetPassport(c *fiber.Ctx, passport interface{}) {
//	c.Locals(CtxPassport, passport)
//}
//
//// GetPassport lấy passport từ context
//func GetPassport(c *fiber.Ctx) interface{} {
//	return c.Locals(CtxPassport)
//}
