package handler

import (
	"net/http"

	"github.com/FelixDu2025/go-practice/internal/service"
)

// TodoHandler 将 HTTP 请求映射到 TodoService。
// TODO: 解析 JSON / PathValue，调用 service，并用 writeJSON / writeError 返回结果。
type TodoHandler struct {
	svc *service.TodoService
}

// NewTodoHandler 创建处理器。
func NewTodoHandler(svc *service.TodoService) *TodoHandler {
	return &TodoHandler{svc: svc}
}

// List TODO: GET /api/todos → 200 + 列表 JSON。
func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: implement List")
}

// Get TODO: GET /api/todos/{id} → 用 r.PathValue("id")；映射 NotFound / InvalidInput。
func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: implement Get")
}

// Create TODO: POST /api/todos → 解码 CreateTodoInput；成功返回 201。
func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: implement Create")
}

// Update TODO: PUT /api/todos/{id} → 解码 UpdateTodoInput（注意 Done 用指针）。
func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: implement Update")
}

// Delete TODO: DELETE /api/todos/{id} → 成功返回 204 No Content。
func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: implement Delete")
}
