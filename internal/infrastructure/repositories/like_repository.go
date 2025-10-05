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

type PostgreSQLLikeRepository struct {
	db *sql.DB
}

func NewLikeRepository(db *sql.DB) repositories.LikeRepository {
	return &PostgreSQLLikeRepository{db: db}
}

func (r *PostgreSQLLikeRepository) Create(like *entities.Like) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO likes (id, user_id, post_id, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, like.ID, like.UserID, like.PostID, like.CreatedAt)
	return err
}

func (r *PostgreSQLLikeRepository) FindByID(id uuid.UUID) (*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, user_id, post_id, created_at FROM likes WHERE id = $1`

	like := &entities.Like{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&like.ID, &like.UserID, &like.PostID, &like.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("like not found")
	}
	return like, err
}

func (r *PostgreSQLLikeRepository) FindByUserAndPost(userID, postID uuid.UUID) (*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, user_id, post_id, created_at FROM likes WHERE user_id = $1 AND post_id = $2`

	like := &entities.Like{}
	err := r.db.QueryRowContext(ctx, query, userID, postID).Scan(&like.ID, &like.UserID, &like.PostID, &like.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("like not found")
	}
	return like, err
}

func (r *PostgreSQLLikeRepository) FindByUserID(userID uuid.UUID, limit, offset int) ([]*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, created_at
		FROM likes
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	return r.queryLikes(ctx, query, userID, limit, offset)
}

func (r *PostgreSQLLikeRepository) FindByPostID(postID uuid.UUID, limit, offset int) ([]*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, created_at
		FROM likes
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	return r.queryLikes(ctx, query, postID, limit, offset)
}

func (r *PostgreSQLLikeRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.db.ExecContext(ctx, "DELETE FROM likes WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("like not found")
	}
	return nil
}

func (r *PostgreSQLLikeRepository) DeleteByUserAndPost(userID, postID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM likes WHERE user_id = $1 AND post_id = $2`
	_, err := r.db.ExecContext(ctx, query, userID, postID)
	return err
}

func (r *PostgreSQLLikeRepository) Exists(id uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := r.db.QueryRowContext(ctx, "SELECT 1 FROM likes WHERE id = $1", id).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (r *PostgreSQLLikeRepository) ExistsByUserAndPost(userID, postID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := r.db.QueryRowContext(ctx, "SELECT 1 FROM likes WHERE user_id = $1 AND post_id = $2", userID, postID).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (r *PostgreSQLLikeRepository) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM likes").Scan(&count)
	return count, err
}

func (r *PostgreSQLLikeRepository) CountByPostID(postID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM likes WHERE post_id = $1", postID).Scan(&count)
	return count, err
}

func (r *PostgreSQLLikeRepository) CountByUserID(userID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM likes WHERE user_id = $1", userID).Scan(&count)
	return count, err
}

// Essential interface methods - shortened
func (r *PostgreSQLLikeRepository) CountByDateRange(start, end time.Time) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM likes WHERE created_at BETWEEN $1 AND $2", start, end).Scan(&count)
	return count, err
}

func (r *PostgreSQLLikeRepository) FindRecentByPost(postID uuid.UUID, limit int) ([]*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, created_at
		FROM likes
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	return r.queryLikes(ctx, query, postID, limit)
}

func (r *PostgreSQLLikeRepository) FindRecentByUser(userID uuid.UUID, limit int) ([]*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, created_at
		FROM likes
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	return r.queryLikes(ctx, query, userID, limit)
}

func (r *PostgreSQLLikeRepository) FindByDateRange(start, end time.Time, limit, offset int) ([]*entities.Like, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, created_at
		FROM likes
		WHERE created_at BETWEEN $1 AND $2
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`

	return r.queryLikes(ctx, query, start, end, limit, offset)
}

func (r *PostgreSQLLikeRepository) GetUsersWhoLiked(postID uuid.UUID, limit, offset int) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT user_id
		FROM likes
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, postID, limit, offset)
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

func (r *PostgreSQLLikeRepository) GetPostsLikedByUser(userID uuid.UUID, limit, offset int) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT post_id
		FROM likes
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postIDs []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		postIDs = append(postIDs, id)
	}

	return postIDs, rows.Err()
}

func (r *PostgreSQLLikeRepository) queryLikes(ctx context.Context, query string, args ...any) ([]*entities.Like, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []*entities.Like
	for rows.Next() {
		like := &entities.Like{}
		err := rows.Scan(&like.ID, &like.UserID, &like.PostID, &like.CreatedAt)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return likes, rows.Err()
}