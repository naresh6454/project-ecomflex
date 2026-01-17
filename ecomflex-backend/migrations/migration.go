package migrations

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/naresh6454/ecomflex-backend/config"
)

// RunMigrations runs database migrations
func RunMigrations(cfg *config.Config) error {
	// Create connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode)
	
	log.Println("Running migrations with connection string (sensitive info redacted):", 
		fmt.Sprintf("postgres://%s:****@%s:%s/%s?sslmode=%s",
		cfg.Database.Username, cfg.Database.Host,
		cfg.Database.Port, cfg.Database.DBName, cfg.Database.SSLMode))
	
	// Create migrate instance
	m, err := migrate.New("file://migrations", connStr)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	
	// Run migrations
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to run - database schema is up to date")
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	
	log.Println("Migrations completed successfully")
	
	return nil
}