package middleware

import (
	"log/slog"
	"net/http"
)

// RequestLogger 记录每个 HTTP 请求的方法、路径、状态码与耗时。
// TODO:
//  1. 包装 ResponseWriter，捕获 WriteHeader 的状态码（默认 200）
//  2. 在 next.ServeHTTP 前后计时
//  3. 用 slog 输出 method / path / status / duration_ms
//
// 当前 stub 仅透传请求，不写日志。
func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	_ = logger
	return func(next http.Handler) http.Handler {
		return next
	}
}
