package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"github.com/labstack/echo"
	"log"
	"poc/web/api"

	userRepository "poc/internal/user/repository"
	userUsecase "poc/internal/user/usecase"
)

type CustomValidator struct {
	validator *validator.Validate
}

func main() {
	//TODO: Implement viper or godotenv for env vars

	sqlDB, gormDB, err := ConnectDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	err = RunMigrations(sqlDB)
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := userRepository.NewGormUserRepository(gormDB)
	userUsec := userUsecase.NewUserUsecase(userRepo)

	e := echo.New()
	e.Validator = buildCustomValidator()
	v1 := e.Group("/v1")
	{
		api.NewUserHandler(v1, userUsec)
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
