package main

import (
	"log"
	"poc/internal/configuration"
)

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
