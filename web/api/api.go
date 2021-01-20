package api

import (
	"github.com/ZupIT/charlescd/web/api/middlewares"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.Use(middlewares.Validator)
	AddHealth(router)


	return router.PathPrefix("/api").Subrouter()
}

func AddHealth(router *mux.Router) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(":)"))
		return
	})
}

func Start(router *mux.Router) {
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Start api on port 8080...")
	log.Fatalln(server.ListenAndServe())
}
