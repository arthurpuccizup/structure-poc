package v1

import (
	"github.com/labstack/echo/v4"
	"net/http"
	userInteractor "poc/internal/use_case/user"
	"poc/web/api/handlers/v1/representation"

	uuidPkg "github.com/google/uuid"
	"poc/internal/logging"
)

func ListUsers(findAllUsers userInteractor.FindAllUsers) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		users, err := findAllUsers.Execute()
		if err != nil {
			logging.LogErrorFromCtx(ctx, err)
			return echoCtx.JSON(http.StatusInternalServerError, err)
		}

		usersResponse := make([]representation.UserResponse, 0)
		for _, user := range users {
			usersResponse = append(usersResponse, representation.FromDomainToResponse(user))
		}

		return echoCtx.JSON(http.StatusOK, usersResponse)
	}
}

func CreateUser(saveUser userInteractor.SaveUser) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		var user representation.UserRequest
		bindErr := echoCtx.Bind(&user)
		if bindErr != nil {
			logging.LogErrorFromCtx(ctx, bindErr)
			return echoCtx.JSON(http.StatusInternalServerError, logging.NewError("Cant parse body", bindErr, nil))
		}

		validationErr := echoCtx.Validate(user)
		if validationErr != nil {
			validationErr = logging.WithOperation(validationErr, "createUser.InputValidation")
			logging.LogErrorFromCtx(ctx, validationErr)
			return echoCtx.JSON(http.StatusInternalServerError, validationErr)
		}

		createdUser, err := saveUser.Execute(user.ToUserDomain())
		if err != nil {
			logging.LogErrorFromCtx(ctx, err)
			return echoCtx.JSON(http.StatusInternalServerError, err)
		}

		return echoCtx.JSON(http.StatusCreated, representation.FromDomainToResponse(createdUser))
	}
}

func FindUserById(findUserById userInteractor.FindUserById) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		uuid, parseErr := uuidPkg.Parse(echoCtx.Param("id"))
		if parseErr != nil {
			logging.LogErrorFromCtx(ctx, parseErr)
			return echoCtx.JSON(http.StatusInternalServerError, logging.NewError("Parse id failed", parseErr, nil))
		}

		user, err := findUserById.Execute(uuid)
		if err != nil {
			logging.LogErrorFromCtx(ctx, err)
			return echoCtx.JSON(http.StatusInternalServerError, err)
		}

		return echoCtx.JSON(http.StatusOK, representation.FromDomainToResponse(user))
	}
}

func UpdateUser(updateUser userInteractor.UpdateUser) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		uuid, parseErr := uuidPkg.Parse(echoCtx.Param("id"))
		if parseErr != nil {
			logging.LogErrorFromCtx(ctx, parseErr)
			return echoCtx.JSON(http.StatusInternalServerError, logging.NewError("Parse id failed", parseErr, nil))
		}

		var user representation.UserRequest
		bindErr := echoCtx.Bind(&user)
		if bindErr != nil {
			logging.LogErrorFromCtx(ctx, bindErr)
			return echoCtx.JSON(http.StatusInternalServerError, logging.NewError("Cant parse body", bindErr, nil))
		}

		validationErr := echoCtx.Validate(user)
		if validationErr != nil {
			validationErr = logging.WithOperation(validationErr, "updateUser.InputValidation")
			logging.LogErrorFromCtx(ctx, validationErr)
			return echoCtx.JSON(http.StatusInternalServerError, logging.NewError("Invalid Fields", validationErr, nil))
		}

		updatedUser, err := updateUser.Execute(uuid, user.ToUserDomain())
		if err != nil {
			logging.LogErrorFromCtx(ctx, bindErr)
			return echoCtx.JSON(http.StatusInternalServerError, err)
		}

		return echoCtx.JSON(http.StatusOK, representation.FromDomainToResponse(updatedUser))
	}
}

func DeleteUSer(deleteUser userInteractor.DeleteUser) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		ctx := echoCtx.Request().Context()
		uuid, parseErr := uuidPkg.Parse(echoCtx.Param("id"))
		if parseErr != nil {
			logging.LogErrorFromCtx(ctx, parseErr)
			return echoCtx.JSON(http.StatusInternalServerError, logging.NewError("Parse id failed", parseErr, nil))
		}

		err := deleteUser.Execute(uuid)
		if err != nil {
			logging.LogErrorFromCtx(ctx, err)
			return echoCtx.JSON(http.StatusInternalServerError, err)
		}

		return echoCtx.JSON(http.StatusNoContent, nil)
	}
}
