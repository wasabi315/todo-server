package repository

import (
	"github.com/gofrs/uuid"
)

type MapRepository struct {
	todos map[uuid.UUID]*Todo
}

func NewMapRepository() (*MapRepository, error) {
	return &MapRepository{
		todos: make(map[uuid.UUID]*Todo),
	}, nil
}
