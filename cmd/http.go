package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"poc/internal/use_case/user"
	handlersV1 "poc/web/api/handlers/v1"
)

type server struct {
	pm   persistenceManager
	echo *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
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
	e.Use(middleware.RequestID())
	e.Validator = buildCustomValidator()

	return e
}

func buildCustomValidator() *CustomValidator {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		log.Fatal(err)
	}

	return &CustomValidator{validator: v}
}

func (s server) registerRoutes() {
	v1 := s.echo.Group("/v1")
	{
		circleHandler := v1.Group("/users")
		{
			circleHandler.GET("", handlersV1.ListUsers(user.NewFindAllUsers(s.pm.userRepository)))
		}
	}
}
