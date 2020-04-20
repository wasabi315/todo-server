package router

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/wasabi315/todo-server/repository"
)

type InternalError struct {
	Err   error
	Stack []byte
}

func (i *InternalError) Error() string {
	return fmt.Sprintf("%s\n%s", i.Err.Error(), i.Stack)
}

var (
	codeStatusMap = map[repository.ErrorCode]int{
		repository.ErrNotFound: http.StatusNotFound,
	}
)

func handleError(err error) error {
	switch v := err.(type) {
	case nil:
		return nil
	case *repository.GeneralError:
		code, ok := codeStatusMap[v.Code]
		if !ok {
			return &InternalError{
				Err:   errors.New(fmt.Sprintf("Unknown error code: %v", v.Code)),
				Stack: debug.Stack(),
			}
		}
		return echo.NewHTTPError(code, v.Error())
	default:
		return &InternalError{err, debug.Stack()}
	}
}

func HTTPErrorHandler(err error, c echo.Context) {
	var (
		code int
		body interface{}
	)

	switch v := err.(type) {
	case nil:
		return
	case *echo.HTTPError:
		if v.Internal != nil {
			if herr, ok := v.Internal.(*echo.HTTPError); ok {
				v = herr
			}
		}
		switch m := v.Message.(type) {
		case string:
			body = echo.Map{"message": m}
		case error:
			body = echo.Map{"message": m.Error()}
		default:
			body = echo.Map{"message": m}
		}
		code = v.Code
	default:
		body = echo.Map{"message": http.StatusText(http.StatusInternalServerError)}
		code = http.StatusInternalServerError
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, body)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
