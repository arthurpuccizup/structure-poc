package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poc/internal/errors"
	"poc/internal/user"
)

type gormRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) user.Repository {
	return gormRepository{db: db}
}

func (r gormRepository) FindAll() ([]User, errors.Error) {
	var users []User

	if res := r.db.Find(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error.Error()).AddOperation("repository.FindAll.Find")
	}

	return users, nil
}

func (r gormRepository) Save(user User) (User, errors.Error) {
	user.ID = uuid.New()
	if res := r.db.Save(&user); res.Error != nil {
		return User{}, errors.New("Save user failed", res.Error.Error()).AddOperation("repository.Save.Save")
	}

	return user, nil
}

func (r gormRepository) GetByID(id uuid.UUID) (User, errors.Error) {
	var user User

	if res := r.db.Model(User{}).Where("id = ?", id).First(&user); res.Error != nil {
		return User{}, errors.New("Find user failed", res.Error.Error()).AddOperation("repository.Save.First")
	}

	return user, nil
}

func (r gormRepository) Update(id uuid.UUID, user User) (User, errors.Error) {
	if res := r.db.Model(User{}).Where("id = ?", id).Updates(&user); res.Error != nil {
		return User{}, errors.New("Update user failed", res.Error.Error()).AddOperation("repository.Update.Updates")
	}

	return user, nil
}

func (r gormRepository) Delete(id uuid.UUID) errors.Error {
	if res := r.db.Delete(User{}, id); res.Error != nil {
		return errors.New("Delete user failed", res.Error.Error()).AddOperation("repository.Delete.Delete")
	}

	return nil
}
