package database

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// Migration represents a database migration
type Migration struct {
	Version string
	Name    string
	SQL     string
}

// Migrator handles database migrations
type Migrator struct {
	db *Database
}

// NewMigrator creates a new migrator
func NewMigrator(db *Database) *Migrator {
	return &Migrator{db: db}
}

// LoadMigrations loads migration files from the embedded filesystem
func (m *Migrator) LoadMigrations() ([]Migration, error) {
	var migrations []Migration

	// Read all SQL files from migrations directory
	entries, err := fs.ReadDir(migrationFiles, "migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		// Extract version from filename (e.g., "001_create_users_table.sql")
		parts := strings.SplitN(entry.Name(), "_", 2)
		if len(parts) < 2 {
			continue
		}

		version := parts[0]
		name := strings.TrimSuffix(parts[1], ".sql")

		// Read SQL content
		sqlPath := filepath.Join("migrations", entry.Name())
		sqlContent, err := fs.ReadFile(migrationFiles, sqlPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read migration file %s: %w", entry.Name(), err)
		}

		migrations = append(migrations, Migration{
			Version: version,
			Name:    name,
			SQL:     string(sqlContent),
		})
	}

	// Sort migrations by version
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

// CreateMigrationsTable creates the migrations tracking table
func (m *Migrator) CreateMigrationsTable(ctx context.Context) error {
	sql := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	_, err := m.db.ExecContext(ctx, sql)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	return nil
}

// GetAppliedMigrations returns a list of already applied migrations
func (m *Migrator) GetAppliedMigrations(ctx context.Context) (map[string]bool, error) {
	sql := `SELECT version FROM schema_migrations;`
	rows, err := m.db.QueryContext(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("failed to query applied migrations: %w", err)
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("failed to scan migration version: %w", err)
		}
		applied[version] = true
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating migration rows: %w", err)
	}

	return applied, nil
}

// Migrate runs all pending migrations
func (m *Migrator) Migrate(ctx context.Context) error {
	// Create migrations table if it doesn't exist
	if err := m.CreateMigrationsTable(ctx); err != nil {
		return err
	}

	// Load all migration files
	migrations, err := m.LoadMigrations()
	if err != nil {
		return err
	}

	// Get applied migrations
	applied, err := m.GetAppliedMigrations(ctx)
	if err != nil {
		return err
	}

	// Run pending migrations
	for _, migration := range migrations {
		if applied[migration.Version] {
			fmt.Printf("Migration %s already applied, skipping\n", migration.Version)
			continue
		}

		fmt.Printf("Applying migration %s: %s\n", migration.Version, migration.Name)

		if err := m.applyMigration(ctx, migration); err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", migration.Version, err)
		}

		fmt.Printf("Successfully applied migration %s\n", migration.Version)
	}

	return nil
}

// applyMigration applies a single migration within a transaction
func (m *Migrator) applyMigration(ctx context.Context, migration Migration) error {
	// Start transaction
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Execute migration SQL
	if _, err := tx.ExecContext(ctx, migration.SQL); err != nil {
		return fmt.Errorf("failed to execute migration SQL: %w", err)
	}

	// Record migration as applied
	insertSQL := `
		INSERT INTO schema_migrations (version, name, applied_at)
		VALUES ($1, $2, $3)
		ON CONFLICT (version) DO NOTHING;
	`

	_, err = tx.ExecContext(ctx, insertSQL, migration.Version, migration.Name, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit migration transaction: %w", err)
	}

	return nil
}

// Rollback rolls back the last migration (if possible)
func (m *Migrator) Rollback(ctx context.Context) error {
	// This is a placeholder for rollback functionality
	// In a production system, you might want to implement down migrations
	return fmt.Errorf("rollback not implemented")
}

// GetMigrationStatus returns the current migration status
func (m *Migrator) GetMigrationStatus(ctx context.Context) ([]MigrationStatus, error) {
	migrations, err := m.LoadMigrations()
	if err != nil {
		return nil, err
	}

	applied, err := m.GetAppliedMigrations(ctx)
	if err != nil {
		return nil, err
	}

	var status []MigrationStatus
	for _, migration := range migrations {
		status = append(status, MigrationStatus{
			Version: migration.Version,
			Name:    migration.Name,
			Applied: applied[migration.Version],
		})
	}

	return status, nil
}

// MigrationStatus represents the status of a migration
type MigrationStatus struct {
	Version string
	Name    string
	Applied bool
}