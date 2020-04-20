package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wasabi315/todo-server/repository"
	"github.com/wasabi315/todo-server/router/params"
)

type Handlers struct {
	Repo repository.Repository
}

func Setup(c *Config) *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = HTTPErrorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := &Handlers{
		Repo: c.Repo,
	}

	api := e.Group("/api")
	{
		apiTodos := api.Group("/todos")
		{
			apiTodos.GET("", h.GetTodos)
			apiTodos.POST("", h.PostTodo)

			apiTodo := apiTodos.Group(param(params.TodoID))
			{
				apiTodo.GET("", h.GetTodoByID)
				apiTodo.PUT("", h.PutTodo)
				apiTodo.DELETE("", h.DeleteTodo)

				apiTodoComplete := apiTodo.Group("/complete")
				{
					apiTodoComplete.PUT("", h.SetTodoComplete)
					apiTodoComplete.DELETE("", h.SetTodoIncomplete)
				}
			}
		}
	}

	return e
}
