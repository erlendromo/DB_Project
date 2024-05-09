package api

import (
	"DB_Project/cmd/api/server"
	"DB_Project/cmd/migrations"
	"DB_Project/internal/config"
	"DB_Project/internal/http/dependencies"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Run() {
	// Load the application configuration
	config := config.NewConfig()

	// Connect to the database and read migrations
	conn := migrations.ConnectToDB(config)
	defer conn.Close()
	m := migrations.MigratePSQL(config)

	// Run migrations and ping the database
	migrations.Guard(m.Up())
	// defer migrations.Guard(m.Down())

	defer func(migrate *migrate.Migrate) {
		if r := recover(); r != nil {
			migrations.Guard(migrate.Down())
		}
	}(m)

	migrations.PingDB(conn)

	// Initialize dependencies
	dependencies.InitDeps(conn)

	// Start the server
	server.StartServer(*config)
}
