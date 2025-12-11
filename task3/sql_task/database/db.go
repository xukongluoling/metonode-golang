package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *gorm.DB

func InitDB() {
	var err error
	connStr := "host=localhost port=54321 user=postgres password=Postgres16 dbname=postgres sslmode=disable"

	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		return
	}
}
