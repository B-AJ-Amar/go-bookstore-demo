package config

import (
	"log"

	"github.com/B-AJ-Amar/go-bookstore-demo/bookstore/models"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// migrate the database schema
	MigrateDB()

	// seed the database with initial data
	// SeedDB()

	log.Println("database Initialized successfully")
}

func GetDB() *gorm.DB {
	return db
}


func MigrateDB() {
	err := db.AutoMigrate(
		&models.Author{},
		&models.Book{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("database migrated successfully")

}