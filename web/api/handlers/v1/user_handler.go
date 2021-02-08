package v1

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	userInteractor "poc/internal/use_case/user"
	"poc/web/api/handlers/v1/representation"

	uuidPkg "github.com/google/uuid"
	"poc/internal/tracking"
)

// @Title List users.
// @Description List all users registered.
// @Success  200  array  representation.UserResponse  "The users registered"
// @Error  200  array  representation.UserResponse  "The users registered"
// @Resource users
// @Route /v1/users [get]
func ListUsers(findAllUsers userInteractor.FindAllUsers) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		logger := ctx.Value(tracking.LoggerFlag).(*zap.SugaredLogger)
		users, err := findAllUsers.Execute(ctx)
		if err != nil {
			logger.Error(err)
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
		logger := c.Request().Context().Value(tracking.LoggerFlag).(*zap.SugaredLogger)
		var user representation.UserRequest
		bindErr := c.Bind(&user)
		if bindErr != nil {
			logger.Error(bindErr)
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Cant parse body", bindErr, nil))
		}

		validationErr := c.Validate(user)
		if validationErr != nil {
			logger.Error(validationErr)
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Invalid Fields", validationErr, nil))
		}

		createdUser, err := saveUser.Execute(user.ToUserDomain())
		if err != nil {
			logger.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, representation.FromDomainToResponse(createdUser))
	}
}

func FindUserById(findUserById userInteractor.FindUserById) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := c.Request().Context().Value(tracking.LoggerFlag).(*zap.SugaredLogger)
		uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
		if parseErr != nil {
			logger.Error(parseErr)
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Parse id failed", parseErr, nil))
		}

		user, err := findUserById.Execute(uuid)
		if err != nil {
			logger.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, representation.FromDomainToResponse(user))
	}
}

func UpdateUser(updateUser userInteractor.UpdateUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := c.Request().Context().Value(tracking.LoggerFlag).(*zap.SugaredLogger)
		uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
		if parseErr != nil {
			logger.Error(parseErr)
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Parse id failed", parseErr, nil))
		}

		var user representation.UserRequest
		bindErr := c.Bind(&user)
		if bindErr != nil {
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Cant parse body", bindErr, nil))
		}

		validationErr := c.Validate(user)
		if validationErr != nil {
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Invalid Fields", validationErr, nil))
		}

		updatedUser, err := updateUser.Execute(uuid, user.ToUserDomain())
		if err != nil {
			logger.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, representation.FromDomainToResponse(updatedUser))
	}
}

func DeleteUSer(deleteUser userInteractor.DeleteUser) echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
		if parseErr != nil {
			return c.JSON(http.StatusInternalServerError, tracking.NewError("Parse id failed", parseErr, nil))
		}

		err := deleteUser.Execute(uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusNoContent, nil)
	}
}
