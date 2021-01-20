package repository

import (
	"github.com/ZupIT/charlescd/internal/errors"
	"github.com/ZupIT/charlescd/internal/models"
	"github.com/ZupIT/charlescd/internal/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type gormRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) user.Repository {
	return gormRepository{db: db}
}

func (r gormRepository) FindAll() ([]models.User, errors.Error) {
	var users []models.User

	if res := r.db.Find(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error.Error()).AddOperation("repository.FindAll.Find")
	}

	return users, nil
}

func (r gormRepository) Save(user models.User) (models.User, errors.Error) {
	user.ID = uuid.New()
	if res := r.db.Save(&user); res.Error != nil {
		return models.User{}, errors.New("Save user failed", res.Error.Error()).AddOperation("repository.Save.Save")
	}

	return user, nil
}

func (r gormRepository) GetByID(id uuid.UUID) (models.User, errors.Error) {
	var user models.User

	if res := r.db.Model(models.User{}).Where("id = ?", id).First(&user); res.Error != nil {
		return models.User{}, errors.New("Find user failed", res.Error.Error()).AddOperation("repository.Save.First")
	}

	return user, nil
}

func (r gormRepository) Update(id uuid.UUID, user models.User) (models.User, errors.Error) {
	if res := r.db.Model(models.User{}).Where("id = ?", id).Updates(&user); res.Error != nil {
		return models.User{}, errors.New("Update user failed", res.Error.Error()).AddOperation("repository.Update.Updates")
	}

	return user, nil
}

func (r gormRepository) Delete(id uuid.UUID) errors.Error {
	if res := r.db.Delete(models.User{}, id); res.Error != nil {
		return errors.New("Delete user failed", res.Error.Error()).AddOperation("repository.Delete.Delete")
	}

	return nil
}
