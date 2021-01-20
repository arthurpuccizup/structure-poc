package user

import (
	"github.com/ZupIT/charlescd/internal/errors"
	"github.com/ZupIT/charlescd/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	FindAll() ([]models.User, errors.Error)
	Save(user models.User) (models.User, errors.Error)
	GetByID(id uuid.UUID) (models.User, errors.Error)
	Update(id uuid.UUID, user models.User) (models.User, errors.Error)
	Delete(id uuid.UUID) errors.Error
}
