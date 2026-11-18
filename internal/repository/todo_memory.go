package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/FelixDu2025/go-practice/internal/model"
)

var (
	// ErrNotFound 表示资源不存在。
	ErrNotFound = errors.New("not found")
	// ErrNotImplemented 表示练习尚未完成实现。
	ErrNotImplemented = errors.New("TODO: repository not implemented")
)

// TodoRepository 定义待办的持久化接口，方便日后换成数据库实现。
type TodoRepository interface {
	List() ([]model.Todo, error)
	Get(id string) (model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Update(todo model.Todo) (model.Todo, error)
	Delete(id string) error
}

// MemoryTodoRepository 使用内存 map 存储，适合练习与测试。
// TODO: 实现并发安全的内存 CRUD（提示：sync.RWMutex + map[string]model.Todo）。
type MemoryTodoRepository struct {
	mu    sync.RWMutex
	items map[string]model.Todo
	now   func() time.Time
	newID func() string
}

// NewMemoryTodoRepository 创建空的内存仓库骨架。
func NewMemoryTodoRepository() *MemoryTodoRepository {
	return &MemoryTodoRepository{
		items: make(map[string]model.Todo),
		now:   time.Now,
		// TODO: 自行实现 ID 生成（可用 crypto/rand 或 uuid 思路）
		newID: func() string { return "" },
	}
}

// List TODO: 返回当前全部待办的副本切片。
func (r *MemoryTodoRepository) List() ([]model.Todo, error) {
	return nil, ErrNotImplemented
}

// Get TODO: 按 id 查找；不存在时返回 ErrNotFound。
func (r *MemoryTodoRepository) Get(id string) (model.Todo, error) {
	return model.Todo{}, ErrNotImplemented
}

// Create TODO: 生成 id，写入 CreatedAt/UpdatedAt，存入 map 并返回。
func (r *MemoryTodoRepository) Create(todo model.Todo) (model.Todo, error) {
	return model.Todo{}, ErrNotImplemented
}

// Update TODO: 保留原 CreatedAt，刷新 UpdatedAt；不存在则 ErrNotFound。
func (r *MemoryTodoRepository) Update(todo model.Todo) (model.Todo, error) {
	return model.Todo{}, ErrNotImplemented
}

// Delete TODO: 删除指定 id；不存在则 ErrNotFound。
func (r *MemoryTodoRepository) Delete(id string) error {
	return ErrNotImplemented
}
