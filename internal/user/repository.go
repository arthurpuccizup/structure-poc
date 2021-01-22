package user

import (
	"github.com/google/uuid"
	"poc/internal/errors"
	"poc/internal/user/repository"
)

type Repository interface {
	FindAll() ([]repository.User, errors.Error)
	Save(user repository.User) (repository.User, errors.Error)
	GetByID(id uuid.UUID) (repository.User, errors.Error)
	Update(id uuid.UUID, user repository.User) (repository.User, errors.Error)
	Delete(id uuid.UUID) errors.Error
}
