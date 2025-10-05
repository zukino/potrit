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

type PostgreSQLConnectionRepository struct {
	db *sql.DB
}

func NewConnectionRepository(db *sql.DB) repositories.ConnectionRepository {
	return &PostgreSQLConnectionRepository{db: db}
}

func (r *PostgreSQLConnectionRepository) Create(connection *entities.Connection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO connections (id, requester_id, addressee_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		connection.ID, connection.RequesterID, connection.AddresseeID,
		connection.Status, connection.CreatedAt, connection.UpdatedAt,
	)

	return err
}

func (r *PostgreSQLConnectionRepository) FindByID(id uuid.UUID) (*entities.Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, requester_id, addressee_id, status, created_at, updated_at
		FROM connections
		WHERE id = $1
	`

	connection := &entities.Connection{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&connection.ID, &connection.RequesterID, &connection.AddresseeID,
		&connection.Status, &connection.CreatedAt, &connection.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("connection not found")
	}
	return connection, err
}

func (r *PostgreSQLConnectionRepository) FindByUsers(requesterID, addresseeID uuid.UUID) (*entities.Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, requester_id, addressee_id, status, created_at, updated_at
		FROM connections
		WHERE (requester_id = $1 AND addressee_id = $2) OR (requester_id = $2 AND addressee_id = $1)
	`

	connection := &entities.Connection{}
	err := r.db.QueryRowContext(ctx, query, requesterID, addresseeID).Scan(
		&connection.ID, &connection.RequesterID, &connection.AddresseeID,
		&connection.Status, &connection.CreatedAt, &connection.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("connection not found")
	}
	return connection, err
}

func (r *PostgreSQLConnectionRepository) FindByUserID(userID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, requester_id, addressee_id, status, created_at, updated_at
		FROM connections
		WHERE (requester_id = $1 OR addressee_id = $1) AND status = $2
		ORDER BY updated_at DESC
		LIMIT $3 OFFSET $4
	`

	return r.queryConnections(ctx, query, userID, status, limit, offset)
}

func (r *PostgreSQLConnectionRepository) Update(connection *entities.Connection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE connections
		SET status = $2, updated_at = $3
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, connection.ID, connection.Status, connection.UpdatedAt)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("connection not found")
	}
	return nil
}

func (r *PostgreSQLConnectionRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM connections WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("connection not found")
	}
	return nil
}

func (r *PostgreSQLConnectionRepository) Exists(id uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := r.db.QueryRowContext(ctx, "SELECT 1 FROM connections WHERE id = $1", id).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (r *PostgreSQLConnectionRepository) ExistsBetweenUsers(requesterID, addresseeID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := r.db.QueryRowContext(ctx, `
		SELECT 1 FROM connections
		WHERE (requester_id = $1 AND addressee_id = $2) OR (requester_id = $2 AND addressee_id = $1)
	`, requesterID, addresseeID).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (r *PostgreSQLConnectionRepository) AreConnected(userID1, userID2 uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := r.db.QueryRowContext(ctx, `
		SELECT 1 FROM connections
		WHERE status = 'accepted' AND ((requester_id = $1 AND addressee_id = $2) OR (requester_id = $2 AND addressee_id = $1))
	`, userID1, userID2).Scan(&exists)

	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (r *PostgreSQLConnectionRepository) GetConnectionStatus(userID1, userID2 uuid.UUID) (entities.ConnectionStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var status entities.ConnectionStatus
	err := r.db.QueryRowContext(ctx, `
		SELECT status FROM connections
		WHERE (requester_id = $1 AND addressee_id = $2) OR (requester_id = $2 AND addressee_id = $1)
		ORDER BY created_at DESC LIMIT 1
	`, userID1, userID2).Scan(&status)

	if err == sql.ErrNoRows {
		return "", fmt.Errorf("no connection found")
	}
	return status, err
}

// Essential interface method implementations - shortened for brevity
func (r *PostgreSQLConnectionRepository) FindByRequesterID(requesterID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.Connection, error) {
	return r.queryConnectionsContext(`WHERE requester_id = $1 AND status = $2`, requesterID, status, limit, offset)
}

func (r *PostgreSQLConnectionRepository) FindByAddresseeID(addresseeID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.Connection, error) {
	return r.queryConnectionsContext(`WHERE addressee_id = $1 AND status = $2`, addresseeID, status, limit, offset)
}

func (r *PostgreSQLConnectionRepository) FindAcceptedConnections(userID uuid.UUID, limit, offset int) ([]*entities.Connection, error) {
	return r.queryConnectionsContext(`WHERE (requester_id = $1 OR addressee_id = $1) AND status = 'accepted'`, userID, limit, offset)
}

func (r *PostgreSQLConnectionRepository) FindPendingRequests(userID uuid.UUID, limit, offset int) ([]*entities.Connection, error) {
	return r.queryConnectionsContext(`WHERE addressee_id = $1 AND status = 'pending'`, userID, limit, offset)
}

func (r *PostgreSQLConnectionRepository) FindSentRequests(userID uuid.UUID, limit, offset int) ([]*entities.Connection, error) {
	return r.queryConnectionsContext(`WHERE requester_id = $1 AND status = 'pending'`, userID, limit, offset)
}

func (r *PostgreSQLConnectionRepository) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM connections").Scan(&count)
	return count, err
}

func (r *PostgreSQLConnectionRepository) CountByStatus(status entities.ConnectionStatus) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM connections WHERE status = $1", status).Scan(&count)
	return count, err
}

func (r *PostgreSQLConnectionRepository) CountByUserID(userID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM connections WHERE requester_id = $1 OR addressee_id = $1", userID).Scan(&count)
	return count, err
}

func (r *PostgreSQLConnectionRepository) GetConnectedUserIDs(userID uuid.UUID) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, `
		SELECT DISTINCT CASE
			WHEN requester_id = $1 THEN addressee_id
			ELSE requester_id
		END as connected_user_id
		FROM connections
		WHERE (requester_id = $1 OR addressee_id = $1) AND status = 'accepted'
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}

	return userIDs, rows.Err()
}

// Helper methods
func (r *PostgreSQLConnectionRepository) queryConnections(ctx context.Context, query string, args ...any) ([]*entities.Connection, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var connections []*entities.Connection
	for rows.Next() {
		conn := &entities.Connection{}
		err := rows.Scan(&conn.ID, &conn.RequesterID, &conn.AddresseeID, &conn.Status, &conn.CreatedAt, &conn.UpdatedAt)
		if err != nil {
			return nil, err
		}
		connections = append(connections, conn)
	}

	return connections, rows.Err()
}

func (r *PostgreSQLConnectionRepository) queryConnectionsContext(whereClause string, args ...any) ([]*entities.Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := fmt.Sprintf(`
		SELECT id, requester_id, addressee_id, status, created_at, updated_at
		FROM connections
		%s
		ORDER BY updated_at DESC
	`, whereClause)

	return r.queryConnections(ctx, query, args...)
}