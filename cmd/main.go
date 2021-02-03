package main

import (
	"log"
	"poc/internal/configuration"
)

// @Version 0.0.1
// @Title POC - API
// @Description POC API, responsible for being the base model for other projects
// @LicenseName Apache 2.0
// @LicenseURL  http://www.apache.org/licenses/LICENSE-2.0
func main() {
	err := configuration.LoadConfigurations()
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
