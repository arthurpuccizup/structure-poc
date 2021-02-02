package user

import (
	"github.com/google/uuid"
	"poc/internal/user/models"
)

type Repository interface {
	FindAll() ([]models.User, error)
	FindAllCustom() ([]models.User, error)
	Save(user models.User) (models.User, error)
	GetByID(id uuid.UUID) (models.User, error)
	Update(id uuid.UUID, user models.User) (models.User, error)
	Delete(id uuid.UUID) error
}
