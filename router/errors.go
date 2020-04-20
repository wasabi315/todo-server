package router

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
)

func HTTPError(code int, err interface{}) error {
	switch v := err.(type) {
	case []interface{}:
		if len(v) == 0 {
			HTTPError(code, nil)
		}
		return HTTPError(code, v[0])
	case nil:
		return echo.NewHTTPError(code)
	case string:
		return echo.NewHTTPError(code, v)
	default:
		return echo.NewHTTPError(code, v)
	}
}

func NotFound(err ...interface{}) error {
	return HTTPError(http.StatusNotFound, err)
}

func BadRequest(err ...interface{}) error {
	return HTTPError(http.StatusBadRequest, err)
}

type InternalError struct {
	Err   error
	Stack []byte
}

func (i *InternalError) Error() string {
	return fmt.Sprintf("%s\n%s", i.Err.Error(), i.Stack)
}

func InternalServerError(err error) error {
	return &InternalError{
		Err:   err,
		Stack: debug.Stack(),
	}
}

func HTTPErrorHandler(err error, c echo.Context) {
	var (
		code int
		body interface{}
	)

	switch e := err.(type) {
	case nil:
		return
	case *echo.HTTPError:
		code = e.Code
		if e.Internal != nil {
			if herr, ok := e.Internal.(*echo.HTTPError); ok {
				e = herr
			}
		}
		body = e.Message
	default:
		code = http.StatusInternalServerError
		body = echo.Map{"message": http.StatusText(http.StatusInternalServerError)}
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
