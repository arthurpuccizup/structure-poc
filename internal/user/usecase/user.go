package usecase

import (
	"encoding/json"
	"io"

	"github.com/ZupIT/charlescd/internal/errors"
	"github.com/ZupIT/charlescd/internal/models"
	"github.com/ZupIT/charlescd/internal/user"
	"github.com/google/uuid"
)

type userUsecase struct {
	userRepository user.Repository
}

func NewUserUsecase(r user.Repository) user.UseCase {
	return userUsecase{
		userRepository: r,
	}
}

func (u userUsecase) Parse(body io.ReadCloser) (models.User, errors.Error) {
	var user models.User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("User parse failed", err.Error())
	}

	return user, nil
}

func (u userUsecase) FindAll() ([]models.User, errors.Error) {
	return u.userRepository.FindAll()
}

func (u userUsecase) Save(user models.User) (models.User, errors.Error) {
	return u.userRepository.Save(user)
}

func (u userUsecase) GetByID(id uuid.UUID) (models.User, errors.Error) {
	return u.userRepository.GetByID(id)
}

func (u userUsecase) Update(id uuid.UUID, user models.User) (models.User, errors.Error) {
	return u.userRepository.Update(id, user)
}

func (u userUsecase) Delete(id uuid.UUID) errors.Error {
	return u.userRepository.Delete(id)
}