package v1

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"poc/internal/user/repository"

	"github.com/google/uuid"
	"poc/internal/errors"
	"poc/internal/user"
)

type UserHandler struct {
	usecase user.UseCase
}

func NewUserHandler(e *echo.Group, u user.UseCase) {
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

func (h UserHandler) list(c echo.Context) error {
	users, err := h.usecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusOK, users)
}

func (h UserHandler) save(c echo.Context) error {
	var user repository.User
	bindErr := c.Bind(&user)
	if bindErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr.Error()).SensitiveError())
	}

	createdUser, err := h.usecase.Save(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusCreated, createdUser)
}

func (h UserHandler) getById(c echo.Context) error {
	uuid, parseErr := uuid.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	user, err := h.usecase.GetByID(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) update(c echo.Context) error {
	uuid, parseErr := uuid.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	var user repository.User
	bindErr := c.Bind(&user)
	if bindErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr.Error()).SensitiveError())
	}

	user, err := h.usecase.Update(uuid, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	return c.JSON(http.StatusOK, user)
}

func (h UserHandler) delete(c echo.Context) error {
	uuid, parseErr := uuid.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr.Error()).SensitiveError())
	}

	err := h.usecase.Delete(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.SensitiveError())
	}

	return c.JSON(http.StatusNoContent, nil)
}
