package unit

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"poc/internal/domain"
	"poc/internal/tracking"
	"poc/internal/use_case/user"
	"poc/tests/unit/fake"
	"testing"
)

type UserTableSuite struct {
	suite.Suite
	findUserById user.FindUserById
}

func (u *UserTableSuite) SetupSuite() {
	u.findUserById = user.NewFindUserById(fake.NewUserRepoFake())
}

func TestSuiteTable(t *testing.T) {
	suite.Run(t, new(UserTableSuite))
}

func (u *UserTableSuite) TestGetByID() {
	tests := map[string]struct {
		input  string
		outPut domain.User
		err    error
	}{
		"ok":    {input: "bf068cc3-d6e4-4d4b-b332-413e397fdac8", outPut: domain.User{Name: "Some Name", Email: "Some Email"}, err: nil},
		"error": {input: "11a6313f-bd56-4f80-be58-ebd3996861c1", outPut: domain.User{}, err: &tracking.CustomError{Title: "Some Title", Detail: "Some Detail", Operations: []string{"findUserByID.Execute"}, Meta: nil}},
	}

	for name, tc := range tests {
		u.Suite.Run(name, func() {
			result, err := u.findUserById.Execute(uuid.MustParse(tc.input))

			assert.Equal(u.Suite.T(), tc.outPut, result)
			assert.Equal(u.Suite.T(), tc.err, err)
		})
	}
}
