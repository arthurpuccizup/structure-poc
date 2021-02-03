package unit

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	models2 "poc/internal/repository/models"
	"poc/internal/use_case/user"
	mocks "poc/tests/unit/mocks/user"
	"testing"
)

type UserSuite struct {
	suite.Suite
	userUC      user.UseCase
	userRepMock *mocks.Repository
}

func (u *UserSuite) SetupSuite() {
	u.userRepMock = new(mocks.Repository)
	u.userUC = user.NewUserUsecase(u.userRepMock)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (u *UserSuite) TestGetByID() {
	u.userRepMock.On("GetByID", mock.Anything).Return(models2.User{ID: uuid.New()}, nil)
	a, err := u.userUC.GetByID(uuid.New())

	require.NotNil(u.T(), a)
	require.Nil(u.T(), err)
}
