package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"log"
	"poc/internal/configuration"
	v1Pkg "poc/web/api/handlers/v1"

	userRepository "poc/internal/user/repository"
	userUseCases "poc/internal/user/usecase"
)

type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	err := configuration.LoadConfigurations()
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, gormDB, err := ConnectDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	err = RunMigrations(sqlDB)
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := userRepository.NewGormUserRepository(gormDB)
	userUsec := userUseCases.NewUserUsecase(userRepo)

	e := echo.New()
	e.Validator = buildCustomValidator()
	v1 := e.Group("/v1")
	{
		v1Pkg.NewUserHandler(v1, userUsec)
	}

	log.Fatalln(e.Start(":8080"))
}

func buildCustomValidator() *CustomValidator {
	v := validator.New()
	err := v.RegisterValidation("notblank", validators.NotBlank)
	if err != nil {
		log.Fatal(err)
	}

	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
