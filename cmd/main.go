package main

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"gorm.io/gorm"
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

	err := loadEnvConfig()
	if err != nil {
		log.Fatal(err)
	}

	gormDB, err := prepareDatabase()
	if err != nil {
		log.Fatal(err)
	}

	http, err := buildHttpHandlers(gormDB)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalln(http.Start(":8080"))
}

func loadEnvConfig() error {
	return configuration.LoadConfigurations()
}

func prepareDatabase() (*gorm.DB, error) {
	sqlDB, gormDB, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}

	err = RunMigrations(sqlDB)
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

func buildHttpHandlers(gormDB *gorm.DB) (*echo.Echo, error) {
	userRepo, err := userRepositoryPkg.NewPostgresUserRepository(gormDB)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Cannot instantiate user repository with error: %s", err.Error()))
	}
	userUseCase := userUseCasesPkg.NewUserUsecase(userRepo)

	e := echo.New()
	e.Validator = buildCustomValidator()
	v1 := e.Group("/v1")
	{
		v1Pkg.NewUserHandler(v1, userUseCase)
	}

	return e, nil
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
