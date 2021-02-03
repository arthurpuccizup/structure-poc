package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	DeletedAt time.Time
}
