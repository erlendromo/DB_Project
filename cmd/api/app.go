package api

import (
	"DB_Project/cmd/api/server"
	"DB_Project/internal/config"
	"DB_Project/internal/datasources/postgresql"
)

func Run() {
	config := config.NewConfig()

	db := postgresql.NewPSQLDatabase(
		config.PSQLConfig.PSQL_host,
		config.PSQLConfig.PSQL_port,
		config.PSQLConfig.PSQL_user,
		config.PSQLConfig.PSQL_password,
		config.PSQLConfig.PSQL_dbname,
	)

	db.CreateDatabase()
	db.Start()
	defer db.Close()
	db.CreateTables()

	server.StartServer(*config)
}
