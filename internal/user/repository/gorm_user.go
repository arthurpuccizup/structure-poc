package repository

import (
	"context"

	"github.com/ZupIT/charlescd/internal/errors"
	"github.com/ZupIT/charlescd/internal/models"
	"github.com/ZupIT/charlescd/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormRepository struct {
	db gorm.DB
}

func NewGormUserRepository() user.Repository {
	return gormRepository{}
}

func (r gormRepository) FindAll(ctx context.Context) ([]models.User, errors.Error) {
	var users []models.User

	if res := r.db.Find(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error.Error()).AddOperation("repository.FindAll.Find")
	}

	return users, nil
}

func (r gormRepository) Save(ctx context.Context, user models.User) (models.User, errors.Error) {
	if res := r.db.Save(user); res.Error != nil {
		return models.User{}, errors.New("Save user failed", res.Error.Error()).AddOperation("repository.Save.Save")
	}

	return user, nil
}

func (r gormRepository) GetByID(ctx context.Context, id uuid.UUID) (models.User, errors.Error) {
	var user models.User

	if res := r.db.Model(models.User{}).Where("id = ?", id).First(&user); res.Error != nil {
		return models.User{}, errors.New("Find user failed", res.Error.Error()).AddOperation("repository.Save.First")
	}

	return user, nil
}

func (r gormRepository) Update(ctx context.Context, id uuid.UUID, user models.User) (models.User, errors.Error) {
	if res := r.db.Model(models.User{}).Where("id = ?", id).Updates(&user); res.Error != nil {
		return models.User{}, errors.New("Update user failed", res.Error.Error()).AddOperation("repository.Update.Updates")
	}

	return user, nil
}

func (r gormRepository) Delete(ctx context.Context, id uuid.UUID) errors.Error {
	if res := r.db.Delete(models.User{}, id); res.Error != nil {
		return errors.New("Delete user failed", res.Error.Error()).AddOperation("repository.Delete.Delete")
	}

	return nil
}
