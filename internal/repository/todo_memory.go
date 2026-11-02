package repository

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"github.com/FelixDu2025/go-practice/internal/model"
)

var (
	// ErrNotFound 表示资源不存在。
	ErrNotFound = errors.New("not found")
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
type MemoryTodoRepository struct {
	mu    sync.RWMutex
	items map[string]model.Todo
	now   func() time.Time
	newID func() string
}

// NewMemoryTodoRepository 创建空的内存仓库。
func NewMemoryTodoRepository() *MemoryTodoRepository {
	return &MemoryTodoRepository{
		items: make(map[string]model.Todo),
		now:   time.Now,
		newID: newID,
	}
}

func newID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		// 极端情况下退回到时间戳，保证练习环境仍可用。
		return hex.EncodeToString([]byte(time.Now().UTC().Format(time.RFC3339Nano)))
	}
	return hex.EncodeToString(b[:])
}

func (r *MemoryTodoRepository) List() ([]model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]model.Todo, 0, len(r.items))
	for _, item := range r.items {
		out = append(out, item)
	}
	return out, nil
}

func (r *MemoryTodoRepository) Get(id string) (model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	item, ok := r.items[id]
	if !ok {
		return model.Todo{}, ErrNotFound
	}
	return item, nil
}

func (r *MemoryTodoRepository) Create(todo model.Todo) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := r.now().UTC()
	todo.ID = r.newID()
	todo.CreatedAt = now
	todo.UpdatedAt = now
	r.items[todo.ID] = todo
	return todo, nil
}

func (r *MemoryTodoRepository) Update(todo model.Todo) (model.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	existing, ok := r.items[todo.ID]
	if !ok {
		return model.Todo{}, ErrNotFound
	}
	todo.CreatedAt = existing.CreatedAt
	todo.UpdatedAt = r.now().UTC()
	r.items[todo.ID] = todo
	return todo, nil
}

func (r *MemoryTodoRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}
	delete(r.items, id)
	return nil
}
