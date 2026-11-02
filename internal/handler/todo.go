package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/FelixDu2025/go-practice/internal/model"
	"github.com/FelixDu2025/go-practice/internal/service"
)

// TodoHandler 将 HTTP 请求映射到 TodoService。
type TodoHandler struct {
	svc *service.TodoService
}

// NewTodoHandler 创建处理器。
func NewTodoHandler(svc *service.TodoService) *TodoHandler {
	return &TodoHandler{svc: svc}
}

func (h *TodoHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.svc.List()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list todos")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	item, err := h.svc.Get(id)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input model.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	item, err := h.svc.Create(input)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var input model.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	item, err := h.svc.Update(id, input)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := h.svc.Delete(id); err != nil {
		h.writeServiceError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TodoHandler) writeServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, service.ErrInvalidInput):
		writeError(w, http.StatusBadRequest, "invalid input")
	case errors.Is(err, service.ErrNotFound):
		writeError(w, http.StatusNotFound, "todo not found")
	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}
