package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	seedFastPowerData(db)
}

func seedFastPowerData(db *sql.DB) {
	stmt, err := db.Prepare(`INSERT INTO fast_power_data (requested_time, received_time, timestamp, manufacturer, manufacturer_serial, generation_active_power, grid_active_power, soc)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	manufacturers := []string{"Alpha", "LG", "Samsung", "Giv"}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(
			time.Now(), time.Now(), time.Now(),
			manufacturers[rand.Intn(len(manufacturers))],
			fmt.Sprintf("BAT%d", rand.Intn(1000)),
			rand.Float64()*10, rand.Float64()*5, rand.Float64()*100,
		)
		if err != nil {
			log.Printf("Failed to insert data: %v", err)
		}
	}

	log.Println("Seed data inserted successfully!")
}
