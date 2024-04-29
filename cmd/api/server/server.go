package server

import (
	"DB_Project/internal/config"
	"DB_Project/internal/http/router"
	"fmt"
	"log"
	"net/http"
)

func StartServer(config config.Config) {
	router := router.NewRouter()

	// TODO Wrap handler with middleware (logger, admin etc)

	log.Printf("Server started on port %s...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), router))
}
