package service

import (
	"errors"
	"testing"

	"github.com/FelixDu2025/go-practice/internal/model"
	"github.com/FelixDu2025/go-practice/internal/repository"
)

func TestTodoServiceCreateValidation(t *testing.T) {
	t.Parallel()

	svc := NewTodoService(repository.NewMemoryTodoRepository())
	if _, err := svc.Create(model.CreateTodoInput{Title: "   "}); !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestTodoServiceCreateAndUpdate(t *testing.T) {
	t.Parallel()

	svc := NewTodoService(repository.NewMemoryTodoRepository())
	created, err := svc.Create(model.CreateTodoInput{Title: " 写笔记 ", Content: "分层架构"})
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if created.Title != "写笔记" {
		t.Fatalf("title should be trimmed, got %q", created.Title)
	}

	done := true
	title := "更新标题"
	updated, err := svc.Update(created.ID, model.UpdateTodoInput{Title: &title, Done: &done})
	if err != nil {
		t.Fatalf("Update: %v", err)
	}
	if !updated.Done || updated.Title != "更新标题" {
		t.Fatalf("unexpected update: %+v", updated)
	}
}

func TestTodoServiceGetNotFound(t *testing.T) {
	t.Parallel()

	svc := NewTodoService(repository.NewMemoryTodoRepository())
	if _, err := svc.Get("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
