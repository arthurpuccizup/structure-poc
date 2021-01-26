package usecase

import (
	"github.com/google/uuid"
	"poc/internal/errors"
	userPkg "poc/internal/user"
	"poc/internal/user/domain"
)

type UseCase interface {
	FindAll() ([]domain.User, errors.Error)
	Save(user domain.User) (domain.User, errors.Error)
	GetByID(id uuid.UUID) (domain.User, errors.Error)
	Update(id uuid.UUID, user domain.User) (domain.User, errors.Error)
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

func (u userUsecase) FindAll() ([]domain.User, errors.Error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return make([]domain.User, 0), err
	}

	domainUsers := make([]domain.User, 0)
	for _, u := range users {
		domainUsers = append(domainUsers, domain.FromUserModel(u))
	}

	return domainUsers, nil
}

func (u userUsecase) Save(user domain.User) (domain.User, errors.Error) {
	savedUser, err := u.userRepository.Save(user.ToUserModel())
	if err != nil {
		return domain.User{}, err
	}

	return domain.FromUserModel(savedUser), nil
}

func (u userUsecase) GetByID(id uuid.UUID) (domain.User, errors.Error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}

	return domain.FromUserModel(user), nil
}

func (u userUsecase) Update(id uuid.UUID, user domain.User) (domain.User, errors.Error) {
	updatedUser, err := u.userRepository.Update(id, user.ToUserModel())
	if err != nil {
		return domain.User{}, err
	}

	return domain.FromUserModel(updatedUser), nil
}

func (u userUsecase) Delete(id uuid.UUID) errors.Error {
	return u.userRepository.Delete(id)
}
