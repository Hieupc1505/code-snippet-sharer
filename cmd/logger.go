package main

//
//import (
//	"context"
//	"fmt"
//	"log/slog"
//	"os"
//	"time"
//)
//
//// DevHandler là custom handler cho chế độ "dev".
//type DevHandler struct {
//	level slog.Level
//	out   *os.File
//}
//
//// NewDevHandler tạo một handler để log với định dạng đơn giản cho môi trường dev.
//func NewDevHandler(out *os.File, level slog.Level) slog.Handler {
//	return &DevHandler{
//		level: level,
//		out:   out,
//	}
//}
//
//// Enabled kiểm tra xem level log có hợp lệ không.
//func (h *DevHandler) Enabled(_ context.Context, level slog.Level) bool {
//	return level >= h.level
//}
//
//// Handle custom cách hiển thị log.
//func (h *DevHandler) Handle(_ context.Context, r slog.Record) error {
//	timestamp := time.Now().Format("15:04:05.000")
//
//	// Format message
//	msg := fmt.Sprintf("[%s] %s: %s", timestamp, r.Level, r.Message)
//
//	// Ghi log ra output
//	fmt.Fprintln(h.out, msg)
//
//	return nil
//}
//
//// WithAttrs và WithGroup không cần xử lý gì trong phiên bản đơn giản.
////func (h *DevHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
////func (h *DevHandler) WithGroup(name string) slog.Handler       { return h }
