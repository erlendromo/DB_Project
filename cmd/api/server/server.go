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
	router := middlewares.LoggerMiddleware(
		middlewares.SessionMiddleware(
			router.NewRouter(),
		),
	)

	log.Printf("Server started on port %s...\n", config.GetApiPort())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.GetApiPort()), router))
}
