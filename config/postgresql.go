package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ragil000/go-restful.git/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection is creating a new connection to database
func Connection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Gagal load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(entities.Book{}, entities.User{})
	return db
}

// CloseConnection is closing a connection from database
func CloseConnection(sql *gorm.DB) {
	db, err := sql.DB()
	if err != nil {
		panic("Failed to close connection database")
	}
	db.Close()
}
