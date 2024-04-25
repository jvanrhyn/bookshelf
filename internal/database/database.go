package database

import (
	"fmt"
	"github.com/jvanrhyn/bookshelf/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

var db *gorm.DB

func InitDatabase() error {
	dsn := buildConnectionString()
	slog.Info("Initializing database", "connection", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		NamingStrategy: PostgresNamingStrategy{},
	})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.Book{}, &model.User{}, &model.Progress{})
	if err != nil {
		return err
	}

	return nil
}

func buildConnectionString() string {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
}
