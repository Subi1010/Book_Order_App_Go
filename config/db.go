package config

import (
	"fmt"
	"log"
	"os"

	"book_order_app/models"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func InitializeDBHandler() *DBHandler {
	db, err := InitPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return &DBHandler{DB: db}
}

// InitPostgresDB initializes a PostgreSQL database connection using GORM
// For dev environment: uses AutoMigrate
// For non-dev environment: uses golang-migrate
func InitPostgresDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOSTNAME, PORT, USERNAME, PASSWORD, DBNAME)

	db, err := gorm.Open(postgresDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Get environment
	env := os.Getenv("APP_ENV")

	// Run auto migration for development environment
	if env == "" || env == "development" || env == "dev" {
		if err := db.AutoMigrate(&models.Book{}, &models.Order{}, &models.User{}); err != nil {
			return nil, fmt.Errorf("error running auto migration: %w", err)
		}
		log.Println("Database auto migration completed successfully (dev environment)")
	} else {
		// Use golang-migrate for non-dev environments
		log.Printf("Running migrations using golang-migrate (environment: %s)", env)
		if err := runMigrations(db, DBNAME); err != nil {
			return nil, fmt.Errorf("error running migrations: %w", err)
		}
		log.Println("Database migrations completed successfully")
	}

	return db, nil
}

// runMigrations runs database migrations using golang-migrate
func runMigrations(gormDB *gorm.DB, dbname string) error {
	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error getting sql.DB from gorm: %w", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		dbname,
		driver,
	)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}
