package repository

import (
	"github.com/gchaincl/dotsql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"poc/internal/errors"
	userPkg "poc/internal/user"
	"poc/internal/user/models"
)

type postgresRepository struct {
	db     *gorm.DB
	dotSql *dotsql.DotSql
}

func NewPostgresUserRepository(db *gorm.DB) (userPkg.Repository, error) {
	dotSql, err := dotsql.LoadFromFile("./internal/user/repository/queries.sql")
	if err != nil {
		return nil, err
	}

	return postgresRepository{db: db, dotSql: dotSql}, nil
}

func (r postgresRepository) FindAll() ([]models.User, errors.Error) {
	var users []models.User

	if res := r.db.Find(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error.Error()).AddOperation("repository.FindAll.Find")
	}

	return users, nil
}

func (r postgresRepository) FindAllCustom() ([]models.User, errors.Error) {
	var users []models.User

	if res := r.db.Raw(r.dotSql.QueryMap()["find-all-custom"]).Scan(&users); res.Error != nil {
		return nil, errors.New("Find all users failed", res.Error.Error()).AddOperation("repository.FindAll.Find")
	}

	return users, nil
}

func (r postgresRepository) Save(user models.User) (models.User, errors.Error) {
	user.ID = uuid.New()
	if res := r.db.Save(&user); res.Error != nil {
		return models.User{}, errors.New("Save userPkg failed", res.Error.Error()).AddOperation("repository.Save.Save")
	}

	return user, nil
}

func (r postgresRepository) GetByID(id uuid.UUID) (models.User, errors.Error) {
	var user models.User

	if res := r.db.Model(models.User{}).Where("id = ?", id).First(&user); res.Error != nil {
		return models.User{}, errors.New("Find userPkg failed", res.Error.Error()).AddOperation("repository.Save.First")
	}

	return user, nil
}

func (r postgresRepository) Update(id uuid.UUID, user models.User) (models.User, errors.Error) {
	if res := r.db.Model(models.User{}).Where("id = ?", id).Updates(&user); res.Error != nil {
		return models.User{}, errors.New("Update userPkg failed", res.Error.Error()).AddOperation("repository.Update.Updates")
	}

	return user, nil
}

func (r postgresRepository) Delete(id uuid.UUID) errors.Error {
	if res := r.db.Delete(models.User{}, id); res.Error != nil {
		return errors.New("Delete userPkg failed", res.Error.Error()).AddOperation("repository.Delete.Delete")
	}

	return nil
}
