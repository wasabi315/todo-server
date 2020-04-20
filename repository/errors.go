package repository

import (
	"errors"
)

var (
	ErrNilID      = errors.New("Nil ID")
	ErrNotFound   = errors.New("Not Found")
	ErrInvalidArg = errors.New("Argument error")
)
