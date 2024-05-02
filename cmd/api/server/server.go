package server

import (
	"DB_Project/internal/config"
	"DB_Project/internal/http/middlewares"
	"DB_Project/internal/http/router"
	"fmt"
	"log"
	"net/http"
)

func StartServer(config config.Config) {
	router := router.NewRouter()

	// TODO Wrap handler with middleware (logger, admin etc)
	logger := middlewares.NewLogger(router)

	log.Printf("Server started on port %s...\n", config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.AppConfig.Port), logger))
}
