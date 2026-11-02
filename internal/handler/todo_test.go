package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FelixDu2025/go-practice/internal/repository"
	"github.com/FelixDu2025/go-practice/internal/service"
)

func TestTodoHandlerCRUDFlow(t *testing.T) {
	t.Parallel()

	repo := repository.NewMemoryTodoRepository()
	svc := service.NewTodoService(repo)
	router := NewRouter(svc)

	createBody := []byte(`{"title":"练习 CRUD","content":"handler 层"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/api/todos", bytes.NewReader(createBody))
	createRec := httptest.NewRecorder()
	router.ServeHTTP(createRec, createReq)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("create status = %d, body = %s", createRec.Code, createRec.Body.String())
	}

	var created map[string]any
	if err := json.Unmarshal(createRec.Body.Bytes(), &created); err != nil {
		t.Fatalf("decode create response: %v", err)
	}
	id, _ := created["id"].(string)
	if id == "" {
		t.Fatal("expected non-empty id")
	}

	listReq := httptest.NewRequest(http.MethodGet, "/api/todos", nil)
	listRec := httptest.NewRecorder()
	router.ServeHTTP(listRec, listReq)
	if listRec.Code != http.StatusOK {
		t.Fatalf("list status = %d", listRec.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/api/todos/"+id, nil)
	getRec := httptest.NewRecorder()
	router.ServeHTTP(getRec, getReq)
	if getRec.Code != http.StatusOK {
		t.Fatalf("get status = %d", getRec.Code)
	}

	updateBody := []byte(`{"done":true}`)
	updateReq := httptest.NewRequest(http.MethodPut, "/api/todos/"+id, bytes.NewReader(updateBody))
	updateRec := httptest.NewRecorder()
	router.ServeHTTP(updateRec, updateReq)
	if updateRec.Code != http.StatusOK {
		t.Fatalf("update status = %d, body = %s", updateRec.Code, updateRec.Body.String())
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/todos/"+id, nil)
	deleteRec := httptest.NewRecorder()
	router.ServeHTTP(deleteRec, deleteReq)
	if deleteRec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d", deleteRec.Code)
	}
}

func TestHealth(t *testing.T) {
	t.Parallel()

	router := NewRouter(service.NewTodoService(repository.NewMemoryTodoRepository()))
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("health status = %d", rec.Code)
	}
}
