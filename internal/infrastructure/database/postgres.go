package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Config holds the database connection configuration
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// DefaultConfig returns a default database configuration
func DefaultConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "password",
		DBName:   "social_media",
		SSLMode:  "disable",
	}
}

// Database wraps the SQL database connection
type Database struct {
	db     *sql.DB
	config Config
}

// New creates a new database connection
func New(config Config) (*Database, error) {
	if config.Host == "" {
		config = DefaultConfig()
	}

	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{
		db:     db,
		config: config,
	}, nil
}

// GetDB returns the underlying SQL database connection
func (d *Database) GetDB() *sql.DB {
	return d.db
}

// Close closes the database connection
func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// Ping tests the database connection
func (d *Database) Ping() error {
	return d.db.Ping()
}

// Health returns the health status of the database
func (d *Database) Health() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := d.db.PingContext(ctx); err != nil {
		return fmt.Errorf("database health check failed: %w", err)
	}

	return nil
}

// GetStats returns connection pool statistics
func (d *Database) GetStats() sql.DBStats {
	return d.db.Stats()
}

// BeginTx starts a transaction
func (d *Database) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.db.BeginTx(ctx, opts)
}

// ExecContext executes a query without returning any rows
func (d *Database) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows
func (d *Database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row
func (d *Database) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}