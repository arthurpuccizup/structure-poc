package user

import (
	"github.com/google/uuid"
	"poc/internal/domain"
	"poc/internal/repository"
	"poc/internal/tracking"
)

type UpdateUser interface {
	Execute(id uuid.UUID, user domain.User) (domain.User, error)
}

type updateUser struct {
	userRepository repository.UserRepository
}

func NewUpdateUser(r repository.UserRepository) UpdateUser {
	return updateUser{
		userRepository: r,
	}
}

func (u updateUser) Execute(id uuid.UUID, user domain.User) (domain.User, error) {
	updatedUser, err := u.userRepository.Update(id, user)
	if err != nil {
		return domain.User{}, tracking.WithOperation(err, "updateUser.Execute")
	}

	return updatedUser, nil
}
