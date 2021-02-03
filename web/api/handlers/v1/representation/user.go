package representation

import (
	"github.com/google/uuid"
	"poc/internal/domain"
)

type UserRequest struct {
	Name  string `json:"name" validate:"required,notblank"`
	Email string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id" example:"6a49fe7b-6586-420e-973a-86be82a79fc2" description:"The user identifier" type:"string" format:"uuid"`
	Name  string    `json:"name" example:"Fulano da Silva" description:"The user name"`
	Email string    `json:"email" example:"user@email.com" description:"The user email"`
}

func (userRequest UserRequest) ToUserDomain() domain.User {
	return domain.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}
}

func FromDomainToResponse(user domain.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
