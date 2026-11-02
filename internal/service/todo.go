package service

import (
	"errors"
	"strings"

	"github.com/FelixDu2025/go-practice/internal/model"
	"github.com/FelixDu2025/go-practice/internal/repository"
)

var (
	// ErrInvalidInput 表示请求参数不合法。
	ErrInvalidInput = errors.New("invalid input")
	// ErrNotFound 透传仓库层的未找到错误。
	ErrNotFound = repository.ErrNotFound
)

// TodoService 封装待办相关业务规则。
type TodoService struct {
	repo repository.TodoRepository
}

// NewTodoService 创建服务实例。
func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) List() ([]model.Todo, error) {
	return s.repo.List()
}

func (s *TodoService) Get(id string) (model.Todo, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return model.Todo{}, ErrInvalidInput
	}
	return s.repo.Get(id)
}

func (s *TodoService) Create(input model.CreateTodoInput) (model.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return model.Todo{}, ErrInvalidInput
	}
	todo := model.Todo{
		Title:   title,
		Content: strings.TrimSpace(input.Content),
		Done:    false,
	}
	return s.repo.Create(todo)
}

func (s *TodoService) Update(id string, input model.UpdateTodoInput) (model.Todo, error) {
	id = strings.TrimSpace(id)
	if id == "" {
		return model.Todo{}, ErrInvalidInput
	}

	existing, err := s.repo.Get(id)
	if err != nil {
		return model.Todo{}, err
	}

	if input.Title != nil {
		title := strings.TrimSpace(*input.Title)
		if title == "" {
			return model.Todo{}, ErrInvalidInput
		}
		existing.Title = title
	}
	if input.Content != nil {
		existing.Content = strings.TrimSpace(*input.Content)
	}
	if input.Done != nil {
		existing.Done = *input.Done
	}

	return s.repo.Update(existing)
}

func (s *TodoService) Delete(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return ErrInvalidInput
	}
	return s.repo.Delete(id)
}
