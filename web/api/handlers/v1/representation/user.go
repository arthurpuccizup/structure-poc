package representation

import (
	"github.com/google/uuid"
	"poc/internal/domain"
)

type UserRequest struct {
	Name  string `json:"name" validate:"required,notblank" conform:"trim"`
	Email string `json:"email" validate:"omitempty,email" conform:"email"`
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
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
