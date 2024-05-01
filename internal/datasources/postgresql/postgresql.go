package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

func StartDB(host, port, user, password string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)

	db, err := sql.Open(user, psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(CREATE_DB)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == "42P04" { // PostgreSQL error code for "database already exists"
				fmt.Println("Database already exists.")
			} else {
				log.Fatal(err)
			}
		}
	}

	log.Println("Created database.")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to database!")
	return db
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

func CreateDBTables(db *sql.DB) {
	for _, query := range CreateTablesList {
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
