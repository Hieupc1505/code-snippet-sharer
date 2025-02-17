package sessionutil

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"s-coder-snippet-sharder/types"
	"sync"
	"time"
)

var (
	SessionStore *session.Store
	once         sync.Once
)

// Khởi tạo session store (Singleton)
func InitSessionStore() {
	once.Do(func() {
		config := session.Config{
			KeyLookup:      types.SessionKey,
			CookieHTTPOnly: true,
			Expiration:     24 * time.Hour, // Session hết hạn sau 24h
		}
		SessionStore = session.New(config)
	})
}

// Hàm Set lưu giá trị vào session theo key
func Set(c *fiber.Ctx, key string, value interface{}) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return err
	}
	sess.Set(key, value)
	return sess.Save()
}

// Hàm Get lấy giá trị từ session theo key
func Get(c *fiber.Ctx, key string) (interface{}, error) {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return nil, err
	}
	value := sess.Get(key)
	if value == nil {
		return nil, fmt.Errorf("key '%s' not found in session", key)
	}
	return value, nil
}

// Hàm Delete xóa key khỏi session
func Delete(c *fiber.Ctx, key string) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return err
	}
	sess.Delete(key)
	return sess.Save()
}

// Hàm Destroy xóa toàn bộ session của user
func Destroy(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return err
	}
	return sess.Destroy()
}
