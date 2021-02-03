package user

import (
	"github.com/google/uuid"
	"poc/internal/domain"
	"poc/internal/errors"
	"poc/internal/repository"
)

type FindUserById interface {
	Execute(id uuid.UUID) (domain.User, error)
}

type findUserById struct {
	userRepository repository.UserRepository
}

func NewFindUserById(r repository.UserRepository) FindUserById {
	return findUserById{
		userRepository: r,
	}
}

func (u findUserById) Execute(id uuid.UUID) (domain.User, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return domain.User{}, errors.WithOperation(err, "findUserByID.Execute")
	}

	return user, nil
}
