package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type Todo struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Notes     string
	Due       *time.Time
	Completed bool
}
