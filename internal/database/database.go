package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jvanrhyn/bookshelf/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// InitDatabase initializes the database connection and performs automatic migration for the defined models.
//
// It takes no parameters and returns an error if the initialization or migration fails.
//
// Example:
//
//	err := database.InitDatabase()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Database initialized successfully.")
func InitDatabase() error {
	dsn := buildConnectionString()
	slog.Info("Initializing database", "connection", dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{

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

// buildConnectionString constructs a PostgreSQL connection string from environment variables.
//
// It retrieves the values of DB_USER, DB_PASS, DB_NAME, DB_HOST, and DB_PORT from the environment and
// constructs a connection string in the format:
//
//	"host=<host> port=<port> user=<user> password=<password> dbname=<dbname> sslmode=disable"
//
// Example:
//
//	connectionString := database.buildConnectionString()
//	fmt.Println(connectionString)
func buildConnectionString() string {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
}
