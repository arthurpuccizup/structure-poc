package user

import (
	"github.com/google/uuid"
	"poc/internal/observ"
	"poc/internal/repository"
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
		return observ.WithOperation(err, "deleteUser.Execute")
	}

	return nil
}
