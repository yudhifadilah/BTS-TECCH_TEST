package models

import (
	"database/sql"
)

type TodoList struct {
	ID          uint         `gorm:"primaryKey"`
	UserID      uint         `json:"user_id" gorm:"not null"`
	Title       string       `json:"title" gorm:"not null"`
	Description string       `json:"description" gorm:"not null"`
	CreatedAt   sql.NullTime `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Todos       []Todo       `json:"todos" gorm:"foreignKey:TodoListID"`
}

type Todo struct {
	ID         uint         `gorm:"primaryKey"`
	TodoListID uint         `json:"todo_list_id" gorm:"not null"`
	Task       string       `json:"task" gorm:"not null"`
	Completed  bool         `json:"completed" gorm:"default:false"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
}
