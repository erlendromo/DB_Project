package api

import (
	"DB_Project/cmd/api/server"
	"DB_Project/internal/config"
)

func Run() {
	// Setup middlewares and initializers here

	config := config.NewConfig()

	server.StartServer(*config)
}
