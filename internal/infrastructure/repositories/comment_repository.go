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

type PostgreSQLCommentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) repositories.CommentRepository {
	return &PostgreSQLCommentRepository{db: db}
}

func (r *PostgreSQLCommentRepository) Create(comment *entities.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		INSERT INTO comments (id, user_id, post_id, content, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		comment.ID, comment.UserID, comment.PostID, comment.Content,
		comment.CreatedAt, comment.UpdatedAt,
	)

	return err
}

func (r *PostgreSQLCommentRepository) FindByID(id uuid.UUID) (*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE id = $1
	`

	comment := &entities.Comment{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&comment.ID, &comment.UserID, &comment.PostID, &comment.Content,
		&comment.CreatedAt, &comment.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("comment not found")
	}
	return comment, err
}

func (r *PostgreSQLCommentRepository) FindByPostID(postID uuid.UUID, limit, offset int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at ASC
		LIMIT $2 OFFSET $3
	`

	return r.queryComments(ctx, query, postID, limit, offset)
}

func (r *PostgreSQLCommentRepository) FindByUserID(userID uuid.UUID, limit, offset int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	return r.queryComments(ctx, query, userID, limit, offset)
}

func (r *PostgreSQLCommentRepository) Update(comment *entities.Comment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE comments
		SET content = $2, updated_at = $3
		WHERE id = $1
	`

	result, err := r.db.ExecContext(ctx, query, comment.ID, comment.Content, comment.UpdatedAt)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("comment not found")
	}
	return nil
}

func (r *PostgreSQLCommentRepository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.db.ExecContext(ctx, "DELETE FROM comments WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("comment not found")
	}
	return nil
}

func (r *PostgreSQLCommentRepository) DeleteByPostID(postID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "DELETE FROM comments WHERE post_id = $1", postID)
	return err
}

func (r *PostgreSQLCommentRepository) DeleteByUserID(userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, "DELETE FROM comments WHERE user_id = $1", userID)
	return err
}

func (r *PostgreSQLCommentRepository) Exists(id uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var exists int
	err := r.db.QueryRowContext(ctx, "SELECT 1 FROM comments WHERE id = $1", id).Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

// Essential interface methods
func (r *PostgreSQLCommentRepository) Count() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM comments").Scan(&count)
	return count, err
}

func (r *PostgreSQLCommentRepository) CountByPostID(postID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM comments WHERE post_id = $1", postID).Scan(&count)
	return count, err
}

func (r *PostgreSQLCommentRepository) CountByUserID(userID uuid.UUID) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM comments WHERE user_id = $1", userID).Scan(&count)
	return count, err
}

func (r *PostgreSQLCommentRepository) FindRecentByPost(postID uuid.UUID, limit int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE post_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	return r.queryComments(ctx, query, postID, limit)
}

func (r *PostgreSQLCommentRepository) Search(query string, limit, offset int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchQuery := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE content ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	return r.queryComments(ctx, searchQuery, "%"+query+"%", limit, offset)
}

func (r *PostgreSQLCommentRepository) GetCommentersByPost(postID uuid.UUID, limit, offset int) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT DISTINCT user_id
		FROM comments
		WHERE post_id = $1
		ORDER BY MIN(created_at) ASC
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

// Helper methods
func (r *PostgreSQLCommentRepository) queryComments(ctx context.Context, query string, args ...any) ([]*entities.Comment, error) {
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*entities.Comment
	for rows.Next() {
		comment := &entities.Comment{}
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content,
			&comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, rows.Err()
}

// Additional essential interface methods
func (r *PostgreSQLCommentRepository) FindMultiple(ids []uuid.UUID) ([]*entities.Comment, error) {
	if len(ids) == 0 {
		return []*entities.Comment{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE id = ANY($1)
		ORDER BY created_at DESC
	`

	return r.queryComments(ctx, query, ids)
}

func (r *PostgreSQLCommentRepository) CountByDateRange(start, end time.Time) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM comments WHERE created_at BETWEEN $1 AND $2", start, end).Scan(&count)
	return count, err
}

func (r *PostgreSQLCommentRepository) FindByDateRange(start, end time.Time, limit, offset int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE created_at BETWEEN $1 AND $2
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`

	return r.queryComments(ctx, query, start, end, limit, offset)
}

func (r *PostgreSQLCommentRepository) FindRecentByUser(userID uuid.UUID, limit int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE user_id = $1
		ORDER BY created_at DESC
		LIMIT $2
	`

	return r.queryComments(ctx, query, userID, limit)
}

func (r *PostgreSQLCommentRepository) FindByUserAndPost(userID, postID uuid.UUID, limit, offset int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE user_id = $1 AND post_id = $2
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`

	return r.queryComments(ctx, query, userID, postID, limit, offset)
}

func (r *PostgreSQLCommentRepository) FindByUserIDs(userIDs []uuid.UUID, limit, offset int) ([]*entities.Comment, error) {
	if len(userIDs) == 0 {
		return []*entities.Comment{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE user_id = ANY($1)
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	return r.queryComments(ctx, query, userIDs, limit, offset)
}

func (r *PostgreSQLCommentRepository) SearchInPost(postID uuid.UUID, query string, limit, offset int) ([]*entities.Comment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	searchQuery := `
		SELECT id, user_id, post_id, content, created_at, updated_at
		FROM comments
		WHERE post_id = $1 AND content ILIKE $2
		ORDER BY created_at DESC
		LIMIT $3 OFFSET $4
	`

	return r.queryComments(ctx, searchQuery, postID, "%"+query+"%", limit, offset)
}

func (r *PostgreSQLCommentRepository) GetMostCommentedPosts(limit, offset int) ([]uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		SELECT post_id
		FROM comments
		GROUP BY post_id
		ORDER BY COUNT(*) DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
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