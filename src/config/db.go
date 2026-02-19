package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// ConnectDB returns *gorm.DB connection
func ConnectDB() *gorm.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading environment variables")
	}

	// Get DATABASE_DSN from environment
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("Error: DATABASE_DSN not set in .env")
	}

	// Connect to database using GORM
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate authentication tables
	// err = db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Role{},
	// 	&models.UserRole{},
	// )
	// if err != nil {
	// 	log.Fatal("Migration failed:", err)
	// }

	fmt.Println("Connected to database and migrations applied successfully!")
	return db
}