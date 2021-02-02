package usecase

import (
	"github.com/google/uuid"
	"poc/internal/errors"
	userPkg "poc/internal/user"
	"poc/internal/user/domain"
	"poc/internal/user/models"
)

type UseCase interface {
	FindAll() ([]domain.User, error)
	Save(user domain.User) (domain.User, error)
	GetByID(id uuid.UUID) (domain.User, error)
	Update(id uuid.UUID, user domain.User) (domain.User, error)
	Delete(id uuid.UUID) error
}

type userUsecase struct {
	userRepository userPkg.Repository
}

func NewUserUsecase(r userPkg.Repository) UseCase {
	return userUsecase{
		userRepository: r,
	}
}

func (u userUsecase) FindAll() ([]domain.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return make([]domain.User, 0), errors.WithOperation(err, "UserUseCase.FindAll")
	}

	domainUsers := make([]domain.User, 0)
	for _, u := range users {
		domainUsers = append(domainUsers, domain.User(u))
	}

	return domainUsers, nil
}

func (u userUsecase) Save(user domain.User) (domain.User, error) {
	savedUser, err := u.userRepository.Save(models.User(user))
	if err != nil {
		return domain.User{}, errors.WithOperation(err, "UserUseCase.Save")
	}

	return domain.User(savedUser), nil
}

func (u userUsecase) GetByID(id uuid.UUID) (domain.User, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return domain.User{}, errors.WithOperation(err, "UserUseCase.GetById")
	}

	return domain.User(user), nil
}

func (u userUsecase) Update(id uuid.UUID, user domain.User) (domain.User, error) {
	updatedUser, err := u.userRepository.Update(id, models.User(user))
	if err != nil {
		return domain.User{}, errors.WithOperation(err, "UserUseCase.Update")
	}

	return domain.User(updatedUser), nil
}

func (u userUsecase) Delete(id uuid.UUID) error {
	return u.userRepository.Delete(id)
}
