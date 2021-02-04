package repository

import (
	"github.com/gchaincl/dotsql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poc/internal/domain"
	"poc/internal/errors"
	models "poc/internal/repository/models"
)

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindAllCustom() ([]domain.User, error)
	Create(user domain.User) (domain.User, error)
	GetByID(id uuid.UUID) (domain.User, error)
	Update(id uuid.UUID, user domain.User) (domain.User, error)
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db     *gorm.DB
	dotSql *dotsql.DotSql
}

func NewUserRepository(db *gorm.DB) (UserRepository, error) {
	dotSql, err := dotsql.LoadFromFile("./internal/repository/queries/user_queries.sql")
	if err != nil {
		return nil, err
	}

	return userRepository{db: db, dotSql: dotSql}, nil
}

func (r userRepository) FindAll() ([]domain.User, error) {
	var users []models.User

	if res := r.db.Find(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error, nil, "repository.FindAll.Find")
	}

	usersFound := make([]domain.User, 0)
	for _, u := range users {
		usersFound = append(usersFound, domain.User(u))
	}

	return usersFound, nil
}

func (r userRepository) FindAllCustom() ([]domain.User, error) {
	var users []models.User

	if res := r.db.Raw(r.dotSql.QueryMap()["find-all-custom"]).Scan(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error, nil, "repository.FindAllCustom.Find")
	}

	usersFound := make([]domain.User, 0)
	for _, u := range users {
		usersFound = append(usersFound, domain.User(u))
	}

	return usersFound, nil
}

func (r userRepository) Create(user domain.User) (domain.User, error) {
	user.ID = uuid.New()
	userToSave := models.User(user)
	if res := r.db.Save(&userToSave); res.Error != nil {
		return domain.User{}, errors.New("Save user failed", res.Error, nil, "repository.Create.Save")
	}

	return user, nil
}

func (r userRepository) GetByID(id uuid.UUID) (domain.User, error) {
	var user models.User

	if res := r.db.Model(models.User{}).Where("id = ?", id).First(&user); res.Error != nil {
		return domain.User{}, errors.New("Find user failed", res.Error, nil, "repository.GetById.First")
	}

	return domain.User(user), nil
}

func (r userRepository) Update(id uuid.UUID, user domain.User) (domain.User, error) {
	userToUpdate := models.User(user)
	if res := r.db.Model(models.User{}).Where("id = ?", id).Updates(&userToUpdate); res.Error != nil {
		return domain.User{}, errors.New("Update user failed", res.Error, nil, "repository.Update.Updates")
	}

	return user, nil
}

func (r userRepository) Delete(id uuid.UUID) error {
	if res := r.db.Delete(models.User{}, id); res.Error != nil {
		return errors.New("Delete user failed", res.Error, nil, "repository.Delete.Delete")
	}

	return nil
}
