package v1

import (
	"github.com/labstack/echo"
	"net/http"
	userInteractor "poc/internal/use_case/user"
	"poc/web/api/handlers/v1/representation"

	uuidPkg "github.com/google/uuid"
	"poc/internal/errors"
)

type UserHandler struct {
	findAllUsers userInteractor.FindAllUsers
	saveUser     userInteractor.SaveUser
	updateUser   userInteractor.UpdateUser
	deleteUser   userInteractor.DeleteUser
	findUserById userInteractor.FindUserById
}

func NewUserHandler(e *echo.Group, findAllUsers userInteractor.FindAllUsers, saveUser userInteractor.SaveUser, updateUser userInteractor.UpdateUser,
	deleteUser userInteractor.DeleteUser,
	findUserById userInteractor.FindUserById) {

	handler := UserHandler{
		findAllUsers: findAllUsers,
		saveUser:     saveUser,
		updateUser:   updateUser,
		deleteUser:   deleteUser,
		findUserById: findUserById,
	}

	users := e.Group("/users")
	{
		users.GET("", handler.list)
		users.POST("", handler.save)
		users.GET("/:userId", handler.getById)
		users.PUT("/:userId", handler.update)
		users.DELETE("/:userId", handler.delete)
	}
}

// @Title List users.
// @Description List all users registered.
// @Success  200  array  representation.UserResponse  "The users registered"
// @Error  200  array  representation.UserResponse  "The users registered"
// @Resource users
// @Route /v1/users [get]
func (handler UserHandler) list(c echo.Context) error {
	users, err := handler.findAllUsers.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
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
		return c.JSON(http.StatusInternalServerError, errors.New("Cant parse body", bindErr, nil))
	}

	validationErr := c.Validate(user)
	if validationErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Invalid Fields", validationErr, nil))
	}

	createdUser, err := handler.saveUser.Execute(user.ToUserDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, representation.FromDomainToResponse(createdUser))
}

func (handler UserHandler) getById(c echo.Context) error {
	uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr, nil))
	}

	user, err := handler.findUserById.Execute(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, representation.FromDomainToResponse(user))
}

func (handler UserHandler) update(c echo.Context) error {
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

	updatedUser, err := handler.updateUser.Execute(uuid, user.ToUserDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, representation.FromDomainToResponse(updatedUser))
}

func (handler UserHandler) delete(c echo.Context) error {
	uuid, parseErr := uuidPkg.Parse(c.Param("userId"))
	if parseErr != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("Parse id failed", parseErr, nil))
	}

	err := handler.deleteUser.Execute(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}
