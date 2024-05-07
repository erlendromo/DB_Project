package migrations

import (
	"DB_Project/internal/config"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Guard(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectToDB(dbconfig *config.Config) *sql.DB {
	conn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbconfig.GetDBUser(), dbconfig.GetDBPassword(), dbconfig.GetDBHost(), dbconfig.GetDBPort(), dbconfig.GetDBName(),
	))
	Guard(err)
	PingDB(conn)

	return conn
}

func PingDB(conn *sql.DB) {
	Guard(conn.Ping())
}

func MigratePSQL(dbconfig *config.Config) *migrate.Migrate {
	m, err := migrate.New(
		"file://cmd/migrations", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbconfig.GetDBUser(), dbconfig.GetDBPassword(), dbconfig.GetDBHost(), dbconfig.GetDBPort(), dbconfig.GetDBName(),
		))

	Guard(err)

	return m
}
