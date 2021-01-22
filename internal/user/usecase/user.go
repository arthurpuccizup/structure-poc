package usecase

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"poc/internal/errors"
	userPkg "poc/internal/user"
	"poc/internal/user/repository"
)

type UseCase interface {
	Parse(body io.ReadCloser) (repository.User, errors.Error)
	FindAll() ([]repository.User, errors.Error)
	Save(user repository.User) (repository.User, errors.Error)
	GetByID(id uuid.UUID) (repository.User, errors.Error)
	Update(id uuid.UUID, user repository.User) (repository.User, errors.Error)
	Delete(id uuid.UUID) errors.Error
}

type userUsecase struct {
	userRepository userPkg.Repository
}

func NewUserUsecase(r userPkg.Repository) UseCase {
	return userUsecase{
		userRepository: r,
	}
}

func (u userUsecase) Parse(body io.ReadCloser) (repository.User, errors.Error) {
	var user repository.User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return repository.User{}, errors.New("User parse failed", err.Error())
	}

	return user, nil
}

func (u userUsecase) FindAll() ([]repository.User, errors.Error) {
	return u.userRepository.FindAll()
}

func (u userUsecase) Save(user repository.User) (repository.User, errors.Error) {
	return u.userRepository.Save(user)
}

func (u userUsecase) GetByID(id uuid.UUID) (repository.User, errors.Error) {
	return u.userRepository.GetByID(id)
}

func (u userUsecase) Update(id uuid.UUID, user repository.User) (repository.User, errors.Error) {
	return u.userRepository.Update(id, user)
}

func (u userUsecase) Delete(id uuid.UUID) errors.Error {
	return u.userRepository.Delete(id)
}
