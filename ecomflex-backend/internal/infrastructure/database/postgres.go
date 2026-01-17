package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/naresh6454/ecomflex-backend/config"
	"github.com/naresh6454/ecomflex-backend/internal/domain/entity"
	"github.com/naresh6454/ecomflex-backend/internal/util"
	"github.com/naresh6454/ecomflex-backend/internal/util/logger"
)

// NewPostgresConnection creates a new PostgreSQL connection
func NewPostgresConnection(dbConfig config.DatabaseConfig) (*sqlx.DB, error) {
	// Create connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.SSLMode,
	)

	// Initialize logger
	log := logger.NewLogger("info")
	
	// Try to connect with retries
	var db *sqlx.DB
	var lastErr error
	
	maxRetries := 5
	retryInterval := 2 * time.Second
	
	for i := 0; i < maxRetries; i++ {
		log.Info(fmt.Sprintf("Connecting to database (attempt %d/%d)...", i+1, maxRetries))
		
		sqlDB, err := sql.Open("postgres", dsn)
		if err != nil {
			lastErr = fmt.Errorf("failed to open database connection: %w", err)
			log.Error("Database connection error", lastErr)
			time.Sleep(retryInterval)
			continue
		}
		
		// Convert to sqlx.DB
		db = sqlx.NewDb(sqlDB, "postgres")
		
		// Test the connection
		err = db.Ping()
		if err == nil {
			log.Info("Successfully connected to the database")
			
			// Ensure superadmin exists
			if err := EnsureSuperAdmin(db, log); err != nil {
				log.Error("Failed to ensure superadmin existence", err)
				// Continue anyway, don't fail the application
			}
			
			return db, nil
		}
		
		lastErr = fmt.Errorf("failed to ping database: %w", err)
		log.Error("Database ping error", lastErr)
		db.Close()
		time.Sleep(retryInterval)
	}
	
	// If we get here, all connection attempts failed
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, lastErr)
}

// SuperAdminCredentials contains the hardcoded credentials for the superadmin
var SuperAdminCredentials = struct {
	Email     string
	Password  string
	FullName  string
	TenantID  uuid.UUID
}{
	Email:     "admin@ecomflex.com",
	Password:  "Admin@123", // In production, use a stronger password
	FullName:  "System Administrator",
	TenantID:  uuid.MustParse("00000000-0000-0000-0000-000000000000"), // Default tenant ID
}

// EnsureSuperAdmin checks if a superadmin exists and creates one if not
func EnsureSuperAdmin(db *sqlx.DB, log logger.Logger) error {
	log.Info("Checking if superadmin exists...")
	
	// Check if superadmin already exists
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM users WHERE role = $1", entity.RoleSuperAdmin)
	if err != nil {
		return fmt.Errorf("failed to check if superadmin exists: %w", err)
	}
	
	// If superadmin already exists, return
	if count > 0 {
		log.Info("Superadmin already exists")
		return nil
	}
	
	log.Info("Creating superadmin user...")
	
	// Hash the password
	hashedPassword, err := util.HashPassword(SuperAdminCredentials.Password)
	if err != nil {
		return fmt.Errorf("failed to hash superadmin password: %w", err)
	}
	
	// Create the superadmin user
	now := time.Now()
	superadminID := uuid.New()
	
	// Insert the superadmin directly
	_, err = db.Exec(`
		INSERT INTO users (
			id, tenant_id, email, password_hash, full_name, role, 
			status, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		)
	`, 
		superadminID, 
		SuperAdminCredentials.TenantID, 
		SuperAdminCredentials.Email,
		hashedPassword,
		SuperAdminCredentials.FullName,
		entity.RoleSuperAdmin,
		"active",
		now,
		now,
	)
	
	if err != nil {
		return fmt.Errorf("failed to create superadmin user: %w", err)
	}
	
	log.Info("Superadmin created successfully")
	log.Info(fmt.Sprintf("Email: %s, Password: [REDACTED]", SuperAdminCredentials.Email))
	
	return nil
}