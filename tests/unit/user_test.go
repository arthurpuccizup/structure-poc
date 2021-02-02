package unit

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"poc/internal/user/models"
	userUsecase "poc/internal/user/usecase"
	mocks "poc/tests/unit/mocks/user"
	"testing"
)

type UserSuite struct {
	suite.Suite
	userUC      userUsecase.UseCase
	userRepMock *mocks.Repository
}

func (u *UserSuite) SetupSuite() {
	u.userRepMock = new(mocks.Repository)
	u.userUC = userUsecase.NewUserUsecase(u.userRepMock)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (u *UserSuite) TestGetByID() {
	u.userRepMock.On("GetByID", mock.Anything).Return(models.User{ID: uuid.New()}, nil)
	a, err := u.userUC.GetByID(uuid.New())

	require.NotNil(u.T(), a)
	require.Nil(u.T(), err)
}
