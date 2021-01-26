package configuration

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadConfigurations() error {
	err := godotenv.Load("./resources/.env")
	if err != nil {
		return err
	}

	return nil
}

func Get(name string) string {
	return os.Getenv(name)
}
