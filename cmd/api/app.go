package api

import (
	"DB_Project/cmd/api/server"
	"DB_Project/internal/config"
	"DB_Project/internal/datasources/postgresql"
)

func Run() {
	// Setup middlewares and initializers here

	config := config.NewConfig()

	db := postgresql.StartDB(config.PSQL_host, config.PSQL_port, config.PSQL_user, config.PSQL_password)
	defer postgresql.CloseDB(db)
	postgresql.CreateDBTables(db)

	server.StartServer(*config)
}
