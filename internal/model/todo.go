package model

import "time"

// Todo 表示一条待办 / 笔记。
type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateTodoInput 创建待办时的请求体。
type CreateTodoInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdateTodoInput 更新待办时的请求体。
// Done 使用指针，以便区分「未传」和「显式设为 false」。
type UpdateTodoInput struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
	Done    *bool   `json:"done,omitempty"`
}
