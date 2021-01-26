package v1

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	userUsecase "poc/internal/user/usecase"
	"poc/web/api/handlers/v1/representation"

	uuidPkg "github.com/google/uuid"
	"poc/internal/errors"
)

type UserHandler struct {
	usecase userUsecase.UseCase
}

func NewUserHandler(e *echo.Group, u userUsecase.UseCase) {
	handler := UserHandler{
		usecase: u,
	}

	path := "/users"
	e.GET(path, handler.list)
	e.POST(path, handler.save)
	e.GET(fmt.Sprintf("%s/%s", path, ":userId"), handler.getById)
	e.PUT(fmt.Sprintf("%s/%s", path, ":userId"), handler.update)
	e.DELETE(fmt.Sprintf("%s/%s", path, ":userId"), handler.delete)
}

func (handler UserHandler) list(c echo.Context) error {
	users, err := handler.usecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	usersResponse := make([]representation.UserResponse, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, representation.FromDomainToResponse(user))
	}

	return c.JSON(http.StatusOK, usersResponse)
}

func (handler UserHandler) save(c echo.Context) error {
	var user representation.UserRequest
	bindErr := c.Bind(&user)
	if bindErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr.Error()).SensitiveError())
	}

	validationErr := c.Validate(user)
	if validationErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Invalid Fields", validationErr.Error()).SensitiveError())
	}

	createdUser, err := handler.usecase.Save(user.ToUserDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusCreated, representation.FromDomainToResponse(createdUser))
}

func (handler UserHandler) getById(c echo.Context) error {
	uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	user, err := handler.usecase.GetByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusOK, representation.FromDomainToResponse(user))
}

func (handler UserHandler) update(c echo.Context) error {
	uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	var user representation.UserRequest
	bindErr := c.Bind(&user)
	if bindErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr.Error()).SensitiveError())
	}

	validationErr := c.Validate(user)
	if validationErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Invalid Fields", validationErr.Error()).SensitiveError())
	}

	updatedUser, err := handler.usecase.Update(uuid, user.ToUserDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusOK, representation.FromDomainToResponse(updatedUser))
}

func (handler UserHandler) delete(c echo.Context) error {
	uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	err := handler.usecase.Delete(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusNoContent, nil)
}
