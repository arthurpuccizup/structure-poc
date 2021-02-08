package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/leebenson/conform"
	"poc/internal/use_case/user"
	handlersV1 "poc/web/api/handlers/v1"
	"poc/web/api/middlewares"
)

type server struct {
	pm   persistenceManager
	echo *echo.Echo
}

type customBinder struct{}

type CustomValidator struct {
	validator *validator.Validate
}

func newServer(pm persistenceManager) server {
	return server{
		pm:   pm,
		echo: createEchoInstance(),
	}
}

func (s server) start(port string) error {
	s.registerRoutes()
	return s.echo.Start(fmt.Sprintf(":%s", port))
}

func createEchoInstance() *echo.Echo {
	e := echo.New()
	e.Use(echoMiddleware.RequestID())
	e.Use(middlewares.ContextLogger)
	e.Validator = buildCustomValidator()
	e.Binder = echo.Binder(customBinder{})

	return e
}

func (cb customBinder) Bind(i interface{}, c echo.Context) (err error) {
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != nil {
		return err
	}

	return conform.Strings(i)
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func buildCustomValidator() *CustomValidator {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		return nil
	}

	return &CustomValidator{validator: v}
}

func (s server) registerRoutes() {
	api := s.echo.Group("/api")
	v1 := api.Group("/v1")
	{
		userHandler := v1.Group("/users")
		{
			userHandler.GET("", handlersV1.ListUsers(user.NewFindAllUsers(s.pm.userRepository)))
			userHandler.POST("", handlersV1.CreateUser(user.NewCreateUser(s.pm.userRepository)))
			userHandler.GET("/:id", handlersV1.FindUserById(user.NewFindUserById(s.pm.userRepository)))
			userHandler.PUT("/:id", handlersV1.UpdateUser(user.NewUpdateUser(s.pm.userRepository)))
			userHandler.DELETE("/:id", handlersV1.DeleteUSer(user.NewDeleteUser(s.pm.userRepository)))
		}
	}
}
