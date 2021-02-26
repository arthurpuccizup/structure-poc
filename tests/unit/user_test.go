package unit

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"poc/internal/domain"
	"poc/internal/logging"
	"poc/internal/use_case/user"
	mocks "poc/tests/unit/mocks/repository"
	"testing"
)

type UserSuite struct {
	suite.Suite
	findUserById user.FindUserById
	userRep      *mocks.UserRepository
}

func (u *UserSuite) SetupSuite() {
	u.userRep = new(mocks.UserRepository)
	u.findUserById = user.NewFindUserById(u.userRep)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (u *UserSuite) TestGetByID() {
	u.userRep.On("GetByID", mock.Anything).Return(domain.User{ID: uuid.New()}, nil)
	a, err := u.findUserById.Execute(uuid.New())

	require.NotNil(u.T(), a)
	require.Nil(u.T(), err)
}

func (u *UserSuite) TestErrorGetByID() {
	u.userRep.On("GetByID", mock.Anything).Return(domain.User{}, logging.NewError("error", errors.New("some error"), nil))
	_, err := u.findUserById.Execute(uuid.New())

	require.Error(u.T(), err)
}
