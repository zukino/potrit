package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/potrit/internal/domain/entities"
	"github.com/potrit/internal/domain/repositories"
)

// PostgreSQLUserRepository implements the UserRepository interface for PostgreSQL
type PostgreSQLUserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new PostgreSQL user repository
func NewUserRepository(db *sql.DB) repositories.UserRepository {
	return &PostgreSQLUserRepository{db: db}
}

// Create saves a new user to the repository
func (r *PostgreSQLUserRepository) Create(user *entities.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO users (id, email, password_hash, name, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.PasswordHash, user.Name, user.Status,
		user.CreatedAt, user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// FindByID retrieves a user by their ID
func (r *PostgreSQLUserRepository) FindByID(id uuid.UUID) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, password_hash, name, status, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &entities.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by ID: %w", err)
	}

	return user, nil
}

// FindByEmail retrieves a user by their email address
func (r *PostgreSQLUserRepository) FindByEmail(email string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, password_hash, name, status, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &entities.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}

	return user, nil
}

// FindByName retrieves users by name (partial match)
func (r *PostgreSQLUserRepository) FindByName(name string, limit, offset int) ([]*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, password_hash, name, status, created_at, updated_at
		FROM users
		WHERE name ILIKE $1
		ORDER BY name ASC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to find users by name: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

// Search performs a text search across user names and emails
func (r *PostgreSQLUserRepository) Search(query string, limit, offset int) ([]*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchQuery := `
		SELECT id, email, password_hash, name, status, created_at, updated_at
		FROM users
		WHERE name ILIKE $1 OR email ILIKE $1
		ORDER BY
			CASE WHEN name ILIKE $1 THEN 1 ELSE 2 END,
			name ASC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, searchQuery, "%"+query+"%", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to search users: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

// Update updates an existing user in the repository
func (r *PostgreSQLUserRepository) Update(user *entities.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE users
		SET email = $2, password_hash = $3, name = $4, status = $5, updated_at = $6
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query,
		user.ID, user.Email, user.PasswordHash, user.Name, user.Status, user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Delete removes a user from the repository
func (r *PostgreSQLUserRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Exists checks if a user with the given ID exists
func (r *PostgreSQLUserRepository) Exists(id uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT 1 FROM users WHERE id = $1`

	var exists int
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if user exists: %w", err)
	}

	return true, nil
}

// ExistsByEmail checks if a user with the given email exists
func (r *PostgreSQLUserRepository) ExistsByEmail(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT 1 FROM users WHERE email = $1`

	var exists int
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if user exists by email: %w", err)
	}

	return true, nil
}

// FindMultiple retrieves multiple users by their IDs
func (r *PostgreSQLUserRepository) FindMultiple(ids []uuid.UUID) ([]*entities.User, error) {
	if len(ids) == 0 {
		return []*entities.User{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, password_hash, name, status, created_at, updated_at
		FROM users
		WHERE id = ANY($1)
		ORDER BY name ASC
	`

	rows, err := r.db.QueryContext(ctx, query, ids)
	if err != nil {
		return nil, fmt.Errorf("failed to find multiple users: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

// Count returns the total number of users
func (r *PostgreSQLUserRepository) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(*) FROM users`

	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}

	return count, nil
}

// CountByStatus returns the number of users by status
func (r *PostgreSQLUserRepository) CountByStatus(status entities.UserStatus) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT COUNT(*) FROM users WHERE status = $1`

	var count int
	err := r.db.QueryRowContext(ctx, query, status).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count users by status: %w", err)
	}

	return count, nil
}

// FindActive returns only active users
func (r *PostgreSQLUserRepository) FindActive(limit, offset int) ([]*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, email, password_hash, name, status, created_at, updated_at
		FROM users
		WHERE status = 'active'
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to find active users: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}

// FindByConnectionStatus finds users based on their connection status with a given user
func (r *PostgreSQLUserRepository) FindByConnectionStatus(userID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT u.id, u.email, u.password_hash, u.name, u.status, u.created_at, u.updated_at
		FROM users u
		INNER JOIN connections c ON (
			(c.requester_id = $1 AND c.addressee_id = u.id) OR
			(c.addressee_id = $1 AND c.requester_id = u.id)
		)
		WHERE c.status = $2 AND u.id != $1
		ORDER BY u.name ASC
		LIMIT $3 OFFSET $4
	`

	rows, err := r.db.QueryContext(ctx, query, userID, status, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to find users by connection status: %w", err)
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Status,
			&user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating user rows: %w", err)
	}

	return users, nil
}