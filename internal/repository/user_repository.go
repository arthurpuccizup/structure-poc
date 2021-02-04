package repository

import (
	"github.com/google/uuid"
	"poc/internal/domain"
)

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindAllCustom() ([]domain.User, error)
	Create(user domain.User) (domain.User, error)
	GetByID(id uuid.UUID) (domain.User, error)
	Update(id uuid.UUID, user domain.User) (domain.User, error)
	Delete(id uuid.UUID) error
}
