package main

import (
	"log"

	"github.com/pbtrad/go_batteries/internal/db"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	log.Println("Database is ready to use.")
}
