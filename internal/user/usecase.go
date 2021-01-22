package user

import (
	"github.com/google/uuid"
	"io"
	"poc/internal/errors"
	"poc/internal/models"
)

type UseCase interface {
	Parse(body io.ReadCloser) (models.User, errors.Error)
	FindAll() ([]models.User, errors.Error)
	Save(user models.User) (models.User, errors.Error)
	GetByID(id uuid.UUID) (models.User, errors.Error)
	Update(id uuid.UUID, user models.User) (models.User, errors.Error)
	Delete(id uuid.UUID) errors.Error
}

type UserInput struct {
	Name  string `json:"name" validate:"notblank, required"`
	Email string `json:"email"`
}
