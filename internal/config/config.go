package config

import (
	"DB_Project/internal/constants"
	"log"
	"os"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	port, present := os.LookupEnv(constants.PORT)
	if !present {
		log.Println("Unable to read port from environment.")
		port = "9000"
	}

	return &Config{
		Port: port,
	}
}
