package user

import (
	"github.com/google/uuid"
	"poc/internal/errors"
	"poc/internal/user/models"
)

type Repository interface {
	FindAll() ([]models.User, errors.Error)
	FindAllCustom() ([]models.User, errors.Error)
	Save(user models.User) (models.User, errors.Error)
	GetByID(id uuid.UUID) (models.User, errors.Error)
	Update(id uuid.UUID, user models.User) (models.User, errors.Error)
	Delete(id uuid.UUID) errors.Error
}
