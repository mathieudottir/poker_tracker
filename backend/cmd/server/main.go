package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mathieudottir/poker_tracker/backend/internal/api"
	"github.com/mathieudottir/poker_tracker/backend/internal/config"
	"github.com/mathieudottir/poker_tracker/backend/internal/repository"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	configPath := "config.yaml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = "../config.yaml"
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup logging
	if err := setupLogging(cfg.Logging.File); err != nil {
		log.Printf("Warning: failed to setup file logging: %v", err)
	}

	log.Println("Starting Poker Tracker Backend...")

	// Connect to database
	dsn := cfg.GetDSN()
	repo, err := repository.NewRepository(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	log.Println("Connected to database successfully")

	// Run migrations
	if err := runMigrations(dsn); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed")

	// Setup API
	apiServer := api.NewAPI(repo)
	router := apiServer.SetupRoutes()

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server listening on %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func setupLogging(logFile string) error {
	if logFile == "" {
		return nil
	}

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		return err
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(file)
	return nil
}

func runMigrations(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Read and execute migration files
	migrations := []string{
		"backend/migrations/001_initial_schema.sql",
		"backend/migrations/002_seed_winamax_data.sql",
	}

	for _, migration := range migrations {
		log.Printf("Running migration: %s", migration)

		content, err := os.ReadFile(migration)
		if err != nil {
			return fmt.Errorf("failed to read migration %s: %w", migration, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("failed to execute migration %s: %w", migration, err)
		}
	}

	return nil
}
