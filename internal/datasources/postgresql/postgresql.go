package postgresql

import (
	"DB_Project/internal/datasources/database"
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

type PSQL_Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string

	Database *sql.DB
}

func NewPSQLDatabase(host, port, user, password, dbName string) database.Database {
	return &PSQL_Database{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

func (psqldb *PSQL_Database) CreateDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", psqldb.Host, psqldb.Port, psqldb.User, psqldb.Password)

	db, err := sql.Open(psqldb.User, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(CREATE_DB); err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "42P04" { // PostgreSQL error code for "database already exists"
				log.Println("Database already exists.")
			} else {
				log.Fatal(err)
			}
		}
	} else {
		log.Println("Created database.")
	}

	db.Close()
}

func (psqldb *PSQL_Database) Start() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", psqldb.Host, psqldb.Port, psqldb.User, psqldb.Password, psqldb.DBName)

	db, err := sql.Open(psqldb.User, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database!")
	psqldb.Database = db
}

func (psqldb *PSQL_Database) CreateTables() {
	if psqldb.Database != nil {
		for _, query := range CreateTablesList {
			tx, err := psqldb.Database.Begin()
			if err != nil {
				log.Printf("Unable to establish transaction: %v", err)
			}
			defer tx.Rollback()

			_, err = tx.Exec(query)
			if err != nil {
				log.Fatal(err)
			}

			if err = tx.Commit(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (psqldb *PSQL_Database) Close() {
	if psqldb != nil {
		if err := psqldb.Database.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
