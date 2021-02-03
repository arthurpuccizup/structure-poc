package main

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
	"log"
	"poc/internal/repository"
	repositoryImpl "poc/internal/repository/impl"
	"poc/internal/use_case/user"
	handlersV1 "poc/web/api/handlers/v1"
)

type repositories struct {
	userRepository repository.UserRepository
}

type useCases struct {
	saveUser     user.SaveUser
	updateUser   user.UpdateUser
	deleteUser   user.DeleteUser
	findAllUsers user.FindAllUsers
	findUserById user.FindUserById
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func buildHttpHandlers(gormDB *gorm.DB) (*echo.Echo, error) {
	_, useCases, err := buildDependencies(gormDB)
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Validator = buildCustomValidator()
	v1 := e.Group("/v1")
	{
		handlersV1.NewUserHandler(v1, useCases.findAllUsers, useCases.saveUser, useCases.updateUser, useCases.deleteUser, useCases.findUserById)
	}

	return e, nil
}

func buildDependencies(gorm *gorm.DB) (repositories, useCases, error) {
	repos, err := buildRepositories(gorm)
	if err != nil {
		return repositories{}, useCases{}, err
	}

	return repos, buildUseCases(repos), nil
}

func buildRepositories(gorm *gorm.DB) (repositories, error) {
	userRepo, err := repositoryImpl.NewPostgresUserRepository(gorm)
	if err != nil {
		return repositories{}, errors.New(fmt.Sprintf("Cannot instantiate user repository with error: %s", err.Error()))
	}

	return repositories{
		userRepository: userRepo,
	}, nil
}

func buildUseCases(repos repositories) useCases {
	return useCases{
		saveUser:     user.NewSaveUser(repos.userRepository),
		updateUser:   user.NewUpdateUser(repos.userRepository),
		deleteUser:   user.NewDeleteUser(repos.userRepository),
		findAllUsers: user.NewFindAllUsers(repos.userRepository),
		findUserById: user.NewFindUserById(repos.userRepository),
	}
}

func buildCustomValidator() *CustomValidator {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		log.Fatal(err)
	}

	return &CustomValidator{validator: v}
}
