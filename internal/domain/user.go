package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	DeletedAt time.Time
}
