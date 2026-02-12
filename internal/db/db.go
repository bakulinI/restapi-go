package db

import (
	"CalculatorAppBackend/internal/calculationService"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=postgres password= dbname=postgres port=5432 sslmode=disable"
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&calculationService.Calculation{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db, nil
}
