package fake

import (
	"github.com/google/uuid"
	"poc/internal/domain"
	"poc/internal/tracking"
)

type UserRepoFake struct{}

func NewUserRepoFake() UserRepoFake {
	return UserRepoFake{}
}

func (u UserRepoFake) GetByID(id uuid.UUID) (domain.User, error) {
	switch id.String() {

	case "bf068cc3-d6e4-4d4b-b332-413e397fdac8":
		return domain.User{
			Name:  "Some Name",
			Email: "Some Email",
		}, nil
	case "11a6313f-bd56-4f80-be58-ebd3996861c1":
		return domain.User{}, &tracking.CustomError{Title: "Some Title", Detail: "Some Detail", Operations: nil, Meta: nil}
	}
	return domain.User{}, nil
}

func (u UserRepoFake) FindAll() ([]domain.User, error) {
	return nil, nil
}
func (u UserRepoFake) FindAllCustom() ([]domain.User, error) {
	return nil, nil
}
func (u UserRepoFake) Create(user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
func (u UserRepoFake) Update(id uuid.UUID, user domain.User) (domain.User, error) {
	return domain.User{}, nil
}
func (u UserRepoFake) Delete(id uuid.UUID) error {
	return nil
}
