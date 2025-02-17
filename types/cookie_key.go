package types

// Định nghĩa các key sử dụng trong Fiber Locals
const (
	AccessTokenKey  = "access_token"  // Key cho cookie access token (JWT)
	RefreshTokenKey = "refresh_token" // Key cho cookie refresh token
	CookieToken     = "cookie_token"  // Key cho cookie token xác thực
	CtxPassport     = "passport"      // Lưu thông tin passport (nếu có)
	LoggerCtx       = "logger"        // Lưu logger vào context
)
