package repository

import (
	"testing"

	"github.com/gofrs/uuid"
)

func TestRepositoryImpl_CreateTodo(t *testing.T) {
	t.Parallel()
	repo, assert, _ := setup(t, common)

	todoArg := TodoArg{}
	todo, err := repo.CreateTodo(todoArg)
	if assert.NoError(err) {
		assert.NotNil(todo)
		assert.NotZero(todo.ID)
		assert.NotZero(todo.CreatedAt)
		assert.NotZero(todo.UpdatedAt)
		assert.Equal(todoArg.Name, todo.Name)
		assert.Equal(todoArg.Notes, todo.Notes)
		assert.Equal(todoArg.Due, todo.Due)
		assert.Equal(todoArg.Completed, todo.Completed)
	}
}

func TestRepositoryImpl_ReadTodos(t *testing.T) {
	t.Parallel()
	repo, assert, _ := setup(t, ex)

	r, err := repo.ReadTodos()
	if assert.NoError(err) {
		assert.Empty(r)
	}

	mustMakeTodo(t, repo)

	r, err = repo.ReadTodos()
	if assert.NoError(err) {
		assert.Len(r, 1)
	}
}

func TestRepositoryImpl_ReadTodoByID(t *testing.T) {
	t.Parallel()
	repo, assert, _ := setup(t, common)

	todo := mustMakeTodo(t, repo)

	r, err := repo.ReadTodoByID(todo.ID)
	if assert.NoError(err) {
		assert.NotNil(r)
		assert.Equal(todo.Name, r.Name)
		assert.Equal(todo.Notes, r.Notes)
		assert.Equal(todo.Due, r.Due)
		assert.Equal(todo.Completed, r.Completed)
	}

	_, err = repo.ReadTodoByID(uuid.Nil)
	assert.Error(err)

	_, err = repo.ReadTodoByID(uuid.Must(uuid.NewV4()))
	assert.Error(err)
}

func TestRepositoryImpl_UpdateTodo(t *testing.T) {
	t.Parallel()
	repo, assert, _ := setup(t, common)

	todo := mustMakeTodo(t, repo)
	todoArg := TodoArg{}

	r, err := repo.UpdateTodo(todo.ID, todoArg)
	if assert.NoError(err) {
		assert.NotNil(r)
		assert.Equal(todoArg.Name, r.Name)
		assert.Equal(todoArg.Notes, r.Notes)
		assert.Equal(todoArg.Due, r.Due)
		assert.Equal(todoArg.Completed, r.Completed)
	}

	_, err = repo.UpdateTodo(uuid.Nil, todoArg)
	assert.Error(err)

	_, err = repo.UpdateTodo(uuid.Must(uuid.NewV4()), todoArg)
	assert.Error(err)
}

func TestRepositoryImpl_SetTodoCompleted(t *testing.T) {
	t.Parallel()
	repo, assert, _ := setup(t, common)

	todo := mustMakeTodo(t, repo)

	err := repo.SetTodoCompleted(todo.ID)
	assert.NoError(err)

	err = repo.SetTodoCompleted(uuid.Nil)
	assert.Error(err)

	err = repo.SetTodoCompleted(uuid.Must(uuid.NewV4()))
	assert.Error(err)
}

func TestRepositoryImpl_SetTodoIncompleted(t *testing.T) {
	t.Parallel()
	repo, assert, _ := setup(t, common)

	todo := mustMakeTodo(t, repo)

	err := repo.SetTodoIncompleted(todo.ID)
	assert.NoError(err)

	err = repo.SetTodoIncompleted(uuid.Nil)
	assert.Error(err)

	err = repo.SetTodoIncompleted(uuid.Must(uuid.NewV4()))
	assert.Error(err)
}
