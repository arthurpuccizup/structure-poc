package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leebenson/conform"
	"poc/internal/use_case/user"
	handlersV1 "poc/web/api/handlers/v1"
)

type server struct {
	pm   persistenceManager
	echo *echo.Echo
}

type customBinder struct{}

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

func (s server) registerRoutes() {
	v1 := s.echo.Group("/v1")
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
