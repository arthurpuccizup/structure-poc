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

	persistenceManager, err := prepareDatabase()
	if err != nil {
		log.Fatal(err)
	}

	server, err := newServer(persistenceManager)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatalln(server.start("8080"))
}
