package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"main.go/config"
	"main.go/models"
)

// DB holds the global database instance
var DB Dbinstance

type Dbinstance struct {
	Db *gorm.DB
}

func Connect() {
	// Parse database port from string to integer
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing DB_PORT to uint: %v", err)
	}

	// Create the DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Config("DB_HOST"),
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
		port,
	)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		os.Exit(2)
	}

	log.Println("Successfully connected to the database.")

	// Set logger and run migrations
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Assign the database instance to the global DB variable
	DB = Dbinstance{Db: db}
}
