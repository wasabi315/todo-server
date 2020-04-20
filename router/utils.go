package router

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/wasabi315/todo-server/router/params"
)

func param(p params.Param) string {
	return fmt.Sprintf("/:%v", p)
}

func getParamAsUUID(c echo.Context, p params.Param) uuid.UUID {
	return uuid.FromStringOrNil(c.Param(string(p)))
}
