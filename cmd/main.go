package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/pbtrad/go_batteries/internal/db"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	queries := db.New(database)

	timestampStr := "2024-02-05T12:00:00Z"
	timestamp, err := time.Parse(time.RFC3339, timestampStr)
	if err != nil {
		log.Fatalf("Failed to parse timestamp: %v", err)
	}

	params := db.GetFastPowerDataParams{
		Manufacturer:       "Tesla",
		ManufacturerSerial: "12345ABC",
		Timestamp:          timestamp,
	}

	ctx := context.Background()
	row, err := queries.GetFastPowerData(ctx, params)
	if err == sql.ErrNoRows {
		log.Println("No data found for this query.")
		return
	} else if err != nil {
		log.Fatalf("Query failed: %v", err)
	}

	log.Println("Query executed successfully:", row)
}
