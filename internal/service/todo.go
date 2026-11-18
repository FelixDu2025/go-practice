package service

import (
	"errors"

	"github.com/FelixDu2025/go-practice/internal/model"
	"github.com/FelixDu2025/go-practice/internal/repository"
)

var (
	// ErrInvalidInput 表示请求参数不合法。
	ErrInvalidInput = errors.New("invalid input")
	// ErrNotFound 透传仓库层的未找到错误。
	ErrNotFound = repository.ErrNotFound
	// ErrNotImplemented 表示练习尚未完成实现。
	ErrNotImplemented = errors.New("TODO: service not implemented")
)

// TodoService 封装待办相关业务规则。
// TODO: 在此层做输入校验（trim、title 非空），再调用 repository。
type TodoService struct {
	repo repository.TodoRepository
}

// NewTodoService 创建服务实例。
func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// List TODO: 委托 repository.List。
func (s *TodoService) List() ([]model.Todo, error) {
	return nil, ErrNotImplemented
}

// Get TODO: 校验 id 非空后调用 repo.Get；空 id 返回 ErrInvalidInput。
func (s *TodoService) Get(id string) (model.Todo, error) {
	return model.Todo{}, ErrNotImplemented
}

// Create TODO: trim title/content；title 为空返回 ErrInvalidInput；Done 默认 false。
func (s *TodoService) Create(input model.CreateTodoInput) (model.Todo, error) {
	return model.Todo{}, ErrNotImplemented
}

// Update TODO: 先 Get，再按指针字段局部更新；空 title 视为非法。
func (s *TodoService) Update(id string, input model.UpdateTodoInput) (model.Todo, error) {
	return model.Todo{}, ErrNotImplemented
}

// Delete TODO: 校验 id 后调用 repo.Delete。
func (s *TodoService) Delete(id string) error {
	return ErrNotImplemented
}
