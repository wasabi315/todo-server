package router

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/copier"
	"github.com/wasabi315/todo-server/repository"
)

type TodoResponse struct {
	ID        uuid.UUID  `json:"todoID"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Name      string     `json:"name"`
	Notes     string     `json:"notes"`
	Due       *time.Time `json:"due,omitempty"`
	Completed bool       `json:"completed"`
}

func formatTodo(todo *repository.Todo) TodoResponse {
	var res TodoResponse
	copier.Copy(&res, todo)
	return res
}

func formatTodos(todos []*repository.Todo) []TodoResponse {
	res := make([]TodoResponse, len(todos))
	for i, todo := range todos {
		res[i] = formatTodo(todo)
	}
	return res
}
