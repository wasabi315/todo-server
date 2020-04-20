package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type TodoArg struct {
	Name      string
	Notes     string
	Due       time.Time
	Completed bool
}

type TodoRepository interface {
	CreateTodo(args TodoArg) (*Todo, error)
	ReadTodos() ([]*Todo, error)
	ReadTodoByID(todoID uuid.UUID) (*Todo, error)
	UpdateTodo(todoID uuid.UUID, args TodoArg) (*Todo, error)
	DeleteTodo(todoID uuid.UUID) error
	SetTodoCompleted(todoID uuid.UUID) error
	SetTodoIncompleted(todo uuid.UUID) error
}
