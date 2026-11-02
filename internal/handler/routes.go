package handler

import (
	"net/http"

	"github.com/FelixDu2025/go-practice/internal/middleware"
	"github.com/FelixDu2025/go-practice/internal/service"
)

// NewRouter 组装路由与中间件（使用 Go 1.22+ 标准库 ServeMux）。
func NewRouter(todoSvc *service.TodoService) http.Handler {
	todoHandler := NewTodoHandler(todoSvc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", Health)
	mux.HandleFunc("GET /api/todos", todoHandler.List)
	mux.HandleFunc("POST /api/todos", todoHandler.Create)
	mux.HandleFunc("GET /api/todos/{id}", todoHandler.Get)
	mux.HandleFunc("PUT /api/todos/{id}", todoHandler.Update)
	mux.HandleFunc("DELETE /api/todos/{id}", todoHandler.Delete)

	return middleware.RequestLogger(nil)(mux)
}
