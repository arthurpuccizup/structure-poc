package user

import (
	"context"
	"go.uber.org/zap"
	"poc/internal/domain"
	"poc/internal/observ"
	"poc/internal/repository"
)

type FindAllUsers interface {
	Execute(ctx context.Context) ([]domain.User, error)
}

type findAllUsers struct {
	userRepository repository.UserRepository
}

func NewFindAllUsers(r repository.UserRepository) FindAllUsers {
	return findAllUsers{
		userRepository: r,
	}
}

func (u findAllUsers) Execute(ctx context.Context) ([]domain.User, error) {
	logger := ctx.Value(observ.LoggerFlag).(*zap.SugaredLogger)
	logger.Info("Listing all users...")
	users, err := u.userRepository.FindAll()
	if err != nil {
		return make([]domain.User, 0), observ.WithOperation(err, "findAllUsers.Execute")
	}

	return users, nil
}
