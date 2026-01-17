package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	OAuth    OAuthConfig
	Captcha  CaptchaConfig
	LogLevel string
	Cors     CorsConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret           string
	AccessTokenExp   time.Duration
	RefreshTokenExp  time.Duration
	RefreshTokenSize int
}

// OAuthConfig holds OAuth configuration
type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
}

// CaptchaConfig holds reCAPTCHA configuration
type CaptchaConfig struct {
	SecretKey string
	SiteKey   string
}

// CorsConfig holds CORS configuration
type CorsConfig struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	ExposeHeaders    []string
	AllowCredentials bool
	MaxAge           time.Duration
}

// Load loads configuration from environment variables or config files
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	// Initialize viper for config file
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")

	// Read from config file (if exists)
	if err := v.ReadInConfig(); err != nil {
		// It's okay if config file is not found, we'll use env vars
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Set defaults
	setDefaults(v)

	// Override with environment variables
	v.AutomaticEnv()
	v.SetEnvPrefix("ECOMFLEX")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Build config
	config := &Config{
		Server: ServerConfig{
			Port:         getEnvOrString(v, "server.port"),
			ReadTimeout:  getEnvOrDuration(v, "server.read_timeout"),
			WriteTimeout: getEnvOrDuration(v, "server.write_timeout"),
		},
		Database: DatabaseConfig{
			Host:     getEnvOrString(v, "database.host"),
			Port:     getEnvOrString(v, "database.port"),
			Username: getEnvOrString(v, "database.username"),
			Password: getEnvOrString(v, "database.password"),
			DBName:   getEnvOrString(v, "database.dbname"),
			SSLMode:  getEnvOrString(v, "database.sslmode"),
		},
		Redis: RedisConfig{
			Host:     getEnvOrString(v, "redis.host"),
			Port:     getEnvOrString(v, "redis.port"),
			Password: getEnvOrString(v, "redis.password"),
			DB:       getEnvOrInt(v, "redis.db"),
		},
		JWT: JWTConfig{
			Secret:           getEnvOrString(v, "jwt.secret"),
			AccessTokenExp:   getEnvOrDuration(v, "jwt.access_token_exp"),
			RefreshTokenExp:  getEnvOrDuration(v, "jwt.refresh_token_exp"),
			RefreshTokenSize: getEnvOrInt(v, "jwt.refresh_token_size"),
		},
		OAuth: OAuthConfig{
			GoogleClientID:     getEnvOrString(v, "oauth.google.client_id"),
			GoogleClientSecret: getEnvOrString(v, "oauth.google.client_secret"),
			GoogleRedirectURL:  getEnvOrString(v, "oauth.google.redirect_url"),
		},
		Captcha: CaptchaConfig{
			SecretKey: getEnvOrString(v, "captcha.secret_key"),
			SiteKey:   getEnvOrString(v, "captcha.site_key"),
		},
		LogLevel: getEnvOrString(v, "log_level"),
		Cors: CorsConfig{
			AllowOrigins:     getEnvOrStringSlice(v, "cors.allow_origins"),
			AllowMethods:     getEnvOrStringSlice(v, "cors.allow_methods"),
			AllowHeaders:     getEnvOrStringSlice(v, "cors.allow_headers"),
			ExposeHeaders:    getEnvOrStringSlice(v, "cors.expose_headers"),
			AllowCredentials: getEnvOrBool(v, "cors.allow_credentials"),
			MaxAge:           getEnvOrDuration(v, "cors.max_age"),
		},
	}

	return config, nil
}

func setDefaults(v *viper.Viper) {
	// Server defaults
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.read_timeout", "15s")
	v.SetDefault("server.write_timeout", "15s")

	// Database defaults
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", "5432")
	v.SetDefault("database.username", "postgres")
	v.SetDefault("database.password", "password")
	v.SetDefault("database.dbname", "ecomflex")
	v.SetDefault("database.sslmode", "disable")

	// Redis defaults
	v.SetDefault("redis.host", "localhost")
	v.SetDefault("redis.port", "6379")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)

	// JWT defaults
	v.SetDefault("jwt.secret", "your-secret-key")
	v.SetDefault("jwt.access_token_exp", "15m")
	v.SetDefault("jwt.refresh_token_exp", "168h") // 7 days
	v.SetDefault("jwt.refresh_token_size", 32)

	// Log level default
	v.SetDefault("log_level", "info")

	// CORS defaults
	v.SetDefault("cors.allow_origins", []string{"http://localhost:5173", "https://yourdomain.com"})
	v.SetDefault("cors.allow_methods", []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	v.SetDefault("cors.allow_headers", []string{"Origin", "Content-Type", "Accept", "Authorization"})
	v.SetDefault("cors.expose_headers", []string{"Content-Length"})
	v.SetDefault("cors.allow_credentials", true)
	v.SetDefault("cors.max_age", "12h")
}

// Helper functions to get values from environment or config
func getEnvOrString(v *viper.Viper, key string) string {
	value := os.Getenv("ECOMFLEX_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
	if value != "" {
		return value
	}
	return v.GetString(key)
}

func getEnvOrInt(v *viper.Viper, key string) int {
	value := os.Getenv("ECOMFLEX_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
	if value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return v.GetInt(key)
}

func getEnvOrDuration(v *viper.Viper, key string) time.Duration {
	value := os.Getenv("ECOMFLEX_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
	if value != "" {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return v.GetDuration(key)
}

func getEnvOrBool(v *viper.Viper, key string) bool {
	value := os.Getenv("ECOMFLEX_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
	if value != "" {
		if b, err := strconv.ParseBool(value); err == nil {
			return b
		}
	}
	return v.GetBool(key)
}

func getEnvOrStringSlice(v *viper.Viper, key string) []string {
	value := os.Getenv("ECOMFLEX_" + strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
	if value != "" {
		return strings.Split(value, ",")
	}
	return v.GetStringSlice(key)
}