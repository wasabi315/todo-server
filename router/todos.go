package router

import (
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/wasabi315/todo-server/repository"
	"github.com/wasabi315/todo-server/router/params"
)

type TodoRequest struct {
	Name      string    `json:"name"`
	Notes     string    `json:"notes"`
	Due       time.Time `json:"due"`
	Completed bool      `json:"completed"`
}

func (h *Handlers) GetTodos(c echo.Context) error {
	res, err := h.Repo.ReadTodos()
	if err != nil {
		return handleError(err)
	}
	return c.JSON(http.StatusOK, formatTodos(res))
}

func (h *Handlers) GetTodoByID(c echo.Context) error {
	todoID := getParamAsUUID(c, params.TodoID)
	todo, err := h.Repo.ReadTodoByID(todoID)
	if err != nil {
		return handleError(err)
	}

	return c.JSON(http.StatusOK, formatTodo(todo))
}

func (h *Handlers) PostTodo(c echo.Context) error {
	var req TodoRequest
	if err := c.Bind(&req); err != nil {
		return handleError(err)
	}
	var todoArg repository.TodoArg
	copier.Copy(&todoArg, req)

	todo, err := h.Repo.CreateTodo(todoArg)
	if err != nil {
		return handleError(err)
	}

	return c.JSON(http.StatusOK, formatTodo(todo))
}

func (h *Handlers) PutTodo(c echo.Context) error {
	todoID := getParamAsUUID(c, params.TodoID)
	var req TodoRequest
	if err := c.Bind(&req); err != nil {
		return handleError(err)
	}
	var todoArg repository.TodoArg
	copier.Copy(&todoArg, req)

	todo, err := h.Repo.UpdateTodo(todoID, todoArg)
	if err != nil {
		return handleError(err)
	}

	return c.JSON(http.StatusOK, formatTodo(todo))
}

func (h *Handlers) DeleteTodo(c echo.Context) error {
	todoID := getParamAsUUID(c, params.TodoID)

	err := h.Repo.DeleteTodo(todoID)
	if err != nil {
		return handleError(err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handlers) SetTodoComplete(c echo.Context) error {
	todoID := getParamAsUUID(c, params.TodoID)

	err := h.Repo.SetTodoCompleted(todoID)
	if err != nil {
		return handleError(err)
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handlers) SetTodoIncomplete(c echo.Context) error {
	todoID := getParamAsUUID(c, params.TodoID)

	err := h.Repo.SetTodoIncompleted(todoID)
	if err != nil {
		return handleError(err)
	}

	return c.NoContent(http.StatusOK)
}
