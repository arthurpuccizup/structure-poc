package user

import (
	"io"

	"github.com/ZupIT/charlescd/internal/errors"
	"github.com/ZupIT/charlescd/internal/models"
	"github.com/google/uuid"
)

type UseCase interface {
	Parse(body io.ReadCloser) (models.User, errors.Error)
	FindAll() ([]models.User, errors.Error)
	Save(user models.User) (models.User, errors.Error)
	GetByID(id uuid.UUID) (models.User, errors.Error)
	Update(id uuid.UUID, user models.User) (models.User, errors.Error)
	Delete(id uuid.UUID) errors.Error
}
