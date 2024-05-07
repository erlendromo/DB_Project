package main

import (
	"DB_Project/cmd/api"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	api.Run()
}
