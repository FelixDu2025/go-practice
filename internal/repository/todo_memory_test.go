package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/FelixDu2025/go-practice/internal/model"
)

func TestMemoryTodoRepositoryCRUD(t *testing.T) {
	t.Parallel()

	fixed := time.Date(2026, 3, 15, 10, 0, 0, 0, time.UTC)
	repo := &MemoryTodoRepository{
		items: make(map[string]model.Todo),
		now:   func() time.Time { return fixed },
		newID: func() string { return "todo-1" },
	}

	created, err := repo.Create(model.Todo{Title: "读 Go 文档", Content: "看 net/http"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if created.ID != "todo-1" {
		t.Fatalf("unexpected id: %s", created.ID)
	}
	if !created.CreatedAt.Equal(fixed) {
		t.Fatalf("unexpected created_at: %v", created.CreatedAt)
	}

	got, err := repo.Get("todo-1")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.Title != "读 Go 文档" {
		t.Fatalf("unexpected title: %s", got.Title)
	}

	list, err := repo.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 item, got %d", len(list))
	}

	updatedAt := fixed.Add(time.Minute)
	repo.now = func() time.Time { return updatedAt }
	updated, err := repo.Update(model.Todo{
		ID:      "todo-1",
		Title:   "读 Go 文档",
		Content: "完成 ServeMux 练习",
		Done:    true,
	})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if !updated.Done || !updated.UpdatedAt.Equal(updatedAt) {
		t.Fatalf("unexpected update result: %+v", updated)
	}

	if err := repo.Delete("todo-1"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := repo.Get("todo-1"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestMemoryTodoRepositoryNotFound(t *testing.T) {
	t.Parallel()

	repo := NewMemoryTodoRepository()
	if _, err := repo.Get("missing"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Get: expected ErrNotFound, got %v", err)
	}
	if err := repo.Delete("missing"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Delete: expected ErrNotFound, got %v", err)
	}
}
