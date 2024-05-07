package api

import (
	"DB_Project/cmd/api/server"
	"DB_Project/cmd/migrations"
	"DB_Project/internal/config"

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
	defer migrations.Guard(m.Down())
	migrations.PingDB(conn)

	// Start the server
	server.StartServer(*config)
}
