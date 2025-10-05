package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
	Name        string        `json:"name"`
	Listen      string        `json:"listen"`
	Port        string        `json:"port"`
	APIKey      string        `json:"api_key"`
	PostgreHost string        `json:"postgre_host"`
	PostgrePort string        `json:"postgre_port"`
	Username    string        `json:"postgre_username"`
	Password    string        `json:"postgre_password"`
	Database    string        `json:"postgre_databse"` // Note: typo in original config
	JWTSecret   string        `json:"jwt_secret,omitempty"`
	SSLMode     string        `json:"sslmode,omitempty"`
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	JWTSecret         string `json:"jwt_secret"`
	JWTExpiration     int    `json:"jwt_expiration"` // in minutes
	RefreshExpiration int    `json:"refresh_expiration"` // in minutes
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(configPath string) (*Config, error) {
	// Default configuration
	config := &Config{
		Name:        "potrit-backend",
		Listen:      "127.0.0.1",
		Port:        "9191",
		APIKey:      "12345",
		PostgreHost: "127.0.0.1",
		PostgrePort: "12345",
		Username:    "postgres",
		Password:    "Zuk1n0",
		Database:    "potrit",
		JWTSecret:   "your-secret-key-change-in-production",
		SSLMode:     "disable",
	}

	// Load from file if it exists
	if configPath == "" {
		configPath = "config.json"
	}

	if _, err := os.Stat(configPath); err == nil {
		file, err := os.Open(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open config file: %w", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(config); err != nil {
			return nil, fmt.Errorf("failed to parse config file: %w", err)
		}
	}

	// Override with environment variables if set
	if name := os.Getenv("APP_NAME"); name != "" {
		config.Name = name
	}
	if listen := os.Getenv("APP_LISTEN"); listen != "" {
		config.Listen = listen
	}
	if port := os.Getenv("APP_PORT"); port != "" {
		config.Port = port
	}
	if apiKey := os.Getenv("APP_API_KEY"); apiKey != "" {
		config.APIKey = apiKey
	}
	if dbHost := os.Getenv("DB_HOST"); dbHost != "" {
		config.PostgreHost = dbHost
	}
	if dbPort := os.Getenv("DB_PORT"); dbPort != "" {
		config.PostgrePort = dbPort
	}
	if dbUser := os.Getenv("DB_USER"); dbUser != "" {
		config.Username = dbUser
	}
	if dbPassword := os.Getenv("DB_PASSWORD"); dbPassword != "" {
		config.Password = dbPassword
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.Database = dbName
	}
	if sslMode := os.Getenv("DB_SSLMODE"); sslMode != "" {
		config.SSLMode = sslMode
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		config.JWTSecret = jwtSecret
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

// Validate validates the configuration
func (c *Config) Validate() error {
	if c.Listen == "" {
		return fmt.Errorf("server listen address is required")
	}
	if c.Port == "" {
		return fmt.Errorf("server port is required")
	}
	if c.APIKey == "" {
		return fmt.Errorf("API key is required")
	}
	if c.PostgreHost == "" {
		return fmt.Errorf("database host is required")
	}
	if c.PostgrePort == "" {
		return fmt.Errorf("database port is required")
	}
	if c.Username == "" {
		return fmt.Errorf("database user is required")
	}
	if c.Database == "" {
		return fmt.Errorf("database name is required")
	}
	if c.JWTSecret != "" && len(c.JWTSecret) < 32 {
		return fmt.Errorf("JWT secret must be at least 32 characters")
	}

	return nil
}

// GetServerAddress returns the server address
func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Listen, c.Port)
}

// GetDatabaseConfig returns database configuration in the format expected by the database package
func (c *Config) GetDatabaseConfig() DatabaseConfig {
	sslMode := c.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	return DatabaseConfig{
		Host:     c.PostgreHost,
		Port:     c.PostgrePort,
		User:     c.Username,
		Password: c.Password,
		DBName:   c.Database,
		SSLMode:  sslMode,
	}
}

// GetJWTConfig returns JWT configuration
func (c *Config) GetJWTConfig() JWTConfig {
	expiration := 60 // default 1 hour
	if c.JWTSecret == "" {
		c.JWTSecret = "your-secret-key-change-in-production"
	}

	return JWTConfig{
		Secret:     c.JWTSecret,
		Expiration: expiration,
	}
}

// DatabaseConfig holds database configuration (for compatibility with database package)
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// JWTConfig holds JWT configuration (for compatibility with auth package)
type JWTConfig struct {
	Secret     string `json:"secret"`
	Expiration int    `json:"expiration"` // in minutes
}

