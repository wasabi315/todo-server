package repository

import (
	"github.com/gofrs/uuid"
)

type MapRepository struct {
	todos map[uuid.UUID]*Todo
}

func NewMapRepository() (Repository, error) {
	return &MapRepository{
		todos: make(map[uuid.UUID]*Todo),
	}, nil
}
