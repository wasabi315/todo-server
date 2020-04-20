package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

func (repo *MapRepository) loadTodo(todoID uuid.UUID) (*Todo, error) {
	todo, ok := repo.todos[todoID]
	if !ok {
		return nil, NotFound("A todo with the specified ID was not found.")
	}
	return todo, nil
}

func (repo *MapRepository) CreateTodo(args TodoArg) (*Todo, error) {
	now := time.Now()
	id := uuid.Must(uuid.NewV4())
	todo := &Todo{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      args.Name,
		Notes:     args.Notes,
		Due:       args.Due,
		Completed: args.Completed,
	}
	repo.todos[id] = todo
	return todo, nil
}

func (repo *MapRepository) ReadTodos() ([]*Todo, error) {
	todos := make([]*Todo, 0)
	for _, todo := range repo.todos {
		if (*todo).ID != uuid.Nil {
			todos = append(todos, todo)
		}
	}
	return todos, nil
}

func (repo *MapRepository) ReadTodoByID(todoID uuid.UUID) (*Todo, error) {
	return repo.loadTodo(todoID)
}

func (repo *MapRepository) UpdateTodo(todoID uuid.UUID, args TodoArg) (*Todo, error) {
	todo, err := repo.loadTodo(todoID)
	if err != nil {
		return nil, err
	}
	todo.Name = args.Name
	todo.Notes = args.Notes
	todo.Due = args.Due
	todo.Completed = args.Completed
	todo.UpdatedAt = time.Now()
	return todo, nil
}

func (repo *MapRepository) DeleteTodo(todoID uuid.UUID) error {
	if _, err := repo.loadTodo(todoID); err != nil {
		return err
	}
	delete(repo.todos, todoID)
	return nil
}

func (repo *MapRepository) SetTodoCompleted(todoID uuid.UUID) error {
	todo, err := repo.loadTodo(todoID)
	if err != nil {
		return err
	}
	todo.Completed = true
	return nil
}

func (repo *MapRepository) SetTodoIncompleted(todoID uuid.UUID) error {
	todo, err := repo.loadTodo(todoID)
	if err != nil {
		return err
	}
	todo.Completed = false
	return nil
}
