package unit

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"poc/internal/user/models"
	userUsecase "poc/internal/user/usecase"
	"poc/tests/unit/mocks"
	"testing"
)

type UserSuite struct {
	suite.Suite
	userUC             userUsecase.UseCase
	userRepositoryMock *mocks.UserRepositoryMock
}

func (u *UserSuite) SetupSuite() {
	u.userRepositoryMock = new(mocks.UserRepositoryMock)
	u.userUC = userUsecase.NewUserUsecase(u.userRepositoryMock)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (u *UserSuite) TestGetByID() {
	u.userRepositoryMock.On("GetById", mock.Anything).Return(models.User{ID: uuid.New()}, nil)
	a, err := u.userUC.GetByID(uuid.New())

	require.NotNil(u.T(), a)
	require.Nil(u.T(), err)
}

func (u *UserSuite) TestAnother() {
	a, err := u.userUC.GetByID(uuid.New())

	require.NotNil(u.T(), a)
	require.Nil(u.T(), err)
}
