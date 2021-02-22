package user

import (
	"poc/internal/domain"
	"poc/internal/logging"
	"poc/internal/repository"
)

type FindAllUsers interface {
	Execute() ([]domain.User, error)
}

type findAllUsers struct {
	userRepository repository.UserRepository
}

func NewFindAllUsers(r repository.UserRepository) FindAllUsers {
	return findAllUsers{
		userRepository: r,
	}
}

func (u findAllUsers) Execute() ([]domain.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return make([]domain.User, 0), logging.WithOperation(err, "findAllUsers.Execute")
	}

	return users, nil
}
