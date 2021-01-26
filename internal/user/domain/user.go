package domain

import (
	"github.com/google/uuid"
	"poc/internal/user/models"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	DeletedAt time.Time
}

func (user User) ToUserModel() models.User {
	return models.User{
		Name:  user.Name,
		Email: user.Email,
	}
}

func FromUserModel(user models.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		DeletedAt: user.DeletedAt,
	}
}
