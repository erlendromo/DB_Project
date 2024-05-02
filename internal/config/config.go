package config

import (
	"DB_Project/internal/constants"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig  AppConfig
	PSQLConfig PSQLConfig
}

type AppConfig struct {
	Port string
}

type PSQLConfig struct {
	PSQL_host     string
	PSQL_port     string
	PSQL_user     string
	PSQL_password string
	PSQL_dbname   string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port, present := os.LookupEnv(constants.PORT)
	if !present {
		log.Println("Unable to read port from environment.")
		port = "9000"
	}

	psql_host, present := os.LookupEnv("PSQL_HOST")
	if !present {
		log.Fatal("Unable to read postgresql HOST from environment.")
	}

	psql_port, present := os.LookupEnv("PSQL_PORT")
	if !present {
		log.Fatal("Unable to read postgresql PORT from environment.")
	}

	psql_user, present := os.LookupEnv("PSQL_USER")
	if !present {
		log.Fatal("Unable to read postgresql USER from environment.")
	}

	psql_password, present := os.LookupEnv("PSQL_PASSWORD")
	if !present {
		log.Fatal("Unable to read postgresql PASSWORD from environment.")
	}

	psql_dbname, present := os.LookupEnv("PSQL_DBNAME")
	if !present {
		log.Fatal("Unable to read postgresql DBNAME from environment.")
	}

	return &Config{
		AppConfig: AppConfig{
			Port: port,
		},
		PSQLConfig: PSQLConfig{
			PSQL_host:     psql_host,
			PSQL_port:     psql_port,
			PSQL_user:     psql_user,
			PSQL_password: psql_password,
			PSQL_dbname:   psql_dbname,
		},
	}
}
