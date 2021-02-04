package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
	"log"
	"poc/internal/use_case/user"
	handlersV1 "poc/web/api/handlers/v1"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func buildHttpHandlers(gormDB *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Validator = buildCustomValidator()

	return e
}

func routing(echo *echo.Echo) {
	v1 := echo.Group("/v1")
	{
		circleHandler := v1.Group("/users")
		{
			circleHandler.GET("", handlersV1.ListUsers(user.NewFindAllUsers()))
		}
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
