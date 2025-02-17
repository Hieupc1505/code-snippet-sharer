package apierror

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/html/charset"
	"io"
)

var (
	//errInternal is the default error message that should be used
	//in the case internal server errors
	ErrorInternal   = errors.New("Oops! An error occurred, please try again.")
	ErrUnauthorised = errors.New("You do not have permission to access.")
	ErrForbidden    = errors.New("You are not allowed to perform this action.")
)

// NewToastMessage tạo JSON string từ message và error
func NewToastMessage(err error) string {

	// Tạo object JSON
	toast := map[string]map[string]string{
		"add-toast": {
			"message": err.Error(),
			"type":    "error",
		},
	}
	// Chuyển object thành JSON string
	jsonBytes, _ := json.Marshal(toast)

	return string(jsonBytes)
}

// Chuyển đổi sang UTF-8 nếu chuỗi bị lỗi encoding
// ensureUTF8 chuyển đổi về UTF-8 đúng
func ensureUTF8(s string) string {
	reader := bytes.NewReader([]byte(s))
	decodedReader, err := charset.NewReader(reader, "text/plain; charset=utf-8")
	if err != nil {
		return s
	}
	decoded, _ := io.ReadAll(decodedReader)
	return string(decoded)
}

// ErrorResponse returns a JSON response with the given error message
func ErrorResponse(c *fiber.Ctx, error error) error {
	// Gửi event để hiển thị toast nhưng không thay đổi trang
	c.Set("Content-Type", "application/json; charset=utf-8")
	c.Set("Hx-Trigger", NewToastMessage(error))
	c.Set("HX-Reswap", "none")
	// Trả về status 200 với nội dung rỗng hoặc giữ nguyên trang
	return c.SendStatus(fiber.StatusBadRequest) // hoặc c.SendString("")
}
