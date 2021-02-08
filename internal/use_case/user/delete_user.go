package user

import (
	"github.com/google/uuid"
	"poc/internal/repository"
	"poc/internal/tracking"
)

type DeleteUser interface {
	Execute(id uuid.UUID) error
}

type deleteUser struct {
	userRepository repository.UserRepository
}

func NewDeleteUser(r repository.UserRepository) DeleteUser {
	return deleteUser{
		userRepository: r,
	}
}

func (u deleteUser) Execute(id uuid.UUID) error {
	err := u.userRepository.Delete(id)
	if err != nil {
		return tracking.WithOperation(err, "deleteUser.Execute")
	}

	return nil
}
