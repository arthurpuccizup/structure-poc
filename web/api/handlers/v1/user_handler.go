package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	userInteractor "poc/internal/use_case/user"
	"poc/web/api/handlers/v1/representation"

	uuidPkg "github.com/google/uuid"
	"poc/internal/errors"
)

// @Title List users.
// @Description List all users registered.
// @Success  200  array  representation.UserResponse  "The users registered"
// @Error  200  array  representation.UserResponse  "The users registered"
// @Resource users
// @Route /v1/users [get]
func ListUsers(findAllUsers userInteractor.FindAllUsers) echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := findAllUsers.Execute()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		usersResponse := make([]representation.UserResponse, 0)
		for _, user := range users {
			usersResponse = append(usersResponse, representation.FromDomainToResponse(user))
		}

		return c.JSON(http.StatusOK, usersResponse)
	}
}

func CreateUser(saveUser userInteractor.SaveUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user representation.UserRequest
		bindErr := c.Bind(&user)
		if bindErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr, nil))
		}

		validationErr := c.Validate(user)
		if validationErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Invalid Fields", validationErr, nil))
		}

		createdUser, err := saveUser.Execute(user.ToUserDomain())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, representation.FromDomainToResponse(createdUser))
	}
}

func FindUserById(findUserById userInteractor.FindUserById) echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
		if parseErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr, nil))
		}

		user, err := findUserById.Execute(uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, representation.FromDomainToResponse(user))
	}
}

func UpdateUser(updateUser userInteractor.UpdateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
		if parseErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr, nil))
		}

		var user representation.UserRequest
		bindErr := c.Bind(&user)
		if bindErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr, nil))
		}

		validationErr := c.Validate(user)
		if validationErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Invalid Fields", validationErr, nil))
		}

		updatedUser, err := updateUser.Execute(uuid, user.ToUserDomain())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, representation.FromDomainToResponse(updatedUser))
	}
}

func DeleteUSer(deleteUser userInteractor.DeleteUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
		if parseErr != nil {
			return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr, nil))
		}

		err := deleteUser.Execute(uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
