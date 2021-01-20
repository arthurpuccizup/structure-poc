package main

import (
	"github.com/labstack/echo"
	"log"

	userHttp "github.com/ZupIT/charlescd/internal/user/delivery/http"
	userRepository "github.com/ZupIT/charlescd/internal/user/repository"
	userUsecase "github.com/ZupIT/charlescd/internal/user/usecase"
)

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
	v1 := e.Group("/v1")
	{
		userHttp.NewUserHandler(v1, userUsec)
	}

	log.Fatalln(e.Start(":8080"))
}
