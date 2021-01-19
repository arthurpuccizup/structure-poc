package user

import (
	"context"

	"github.com/ZupIT/charlescd/internal/errors"
	"github.com/ZupIT/charlescd/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	FindAll(ctx context.Context) ([]models.User, errors.Error)
	Save(ctx context.Context, user models.User) (models.User, errors.Error)
	GetByID(ctx context.Context, id uuid.UUID) (models.User, errors.Error)
	Update(ctx context.Context, id uuid.UUID, user models.User) (models.User, errors.Error)
	Delete(ctx context.Context, id uuid.UUID) errors.Error
}
