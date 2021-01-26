package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"log"
	"poc/internal/configuration"
	v1Pkg "poc/web/api/handlers/v1"

	userRepositoryPkg "poc/internal/user/repository"
	userUseCasesPkg "poc/internal/user/usecase"
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

	userRepo, err := userRepositoryPkg.NewPostgresUserRepository(gormDB)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot instantiate user repository with error: %s", err.Error()))
	}
	userUseCase := userUseCasesPkg.NewUserUsecase(userRepo)

	e := echo.New()
	e.Validator = buildCustomValidator()
	v1 := e.Group("/v1")
	{
		v1Pkg.NewUserHandler(v1, userUseCase)
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
