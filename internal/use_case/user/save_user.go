package user

import (
	"poc/internal/domain"
	"poc/internal/errors"
	"poc/internal/repository"
)

type SaveUser interface {
	Execute(user domain.User) (domain.User, error)
}

type saveUser struct {
	userRepository repository.UserRepository
}

func NewSaveUser(r repository.UserRepository) SaveUser {
	return saveUser{
		userRepository: r,
	}
}

func (u saveUser) Execute(user domain.User) (domain.User, error) {
	savedUser, err := u.userRepository.Create(user)
	if err != nil {
		return domain.User{}, errors.WithOperation(err, "saveUser.Execute")
	}

	return savedUser, nil
}
