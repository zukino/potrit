package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/internal/domain/entities"
)

// CommentRepository defines the interface for comment data operations
type CommentRepository interface {
	// Create saves a new comment to the repository
	Create(comment *entities.Comment) error

	// FindByID retrieves a comment by their ID
	FindByID(id uuid.UUID) (*entities.Comment, error)

	// FindByPostID retrieves comments for a specific post
	FindByPostID(postID uuid.UUID, limit, offset int) ([]*entities.Comment, error)

	// FindByUserID retrieves comments made by a specific user
	FindByUserID(userID uuid.UUID, limit, offset int) ([]*entities.Comment, error)

	// FindByUserAndPost retrieves comments by a user on a specific post
	FindByUserAndPost(userID, postID uuid.UUID, limit, offset int) ([]*entities.Comment, error)

	// Update updates an existing comment in the repository
	Update(comment *entities.Comment) error

	// Delete removes a comment from the repository
	Delete(id uuid.UUID) error

	// DeleteByPostID removes all comments for a post
	DeleteByPostID(postID uuid.UUID) error

	// DeleteByUserID removes all comments made by a user
	DeleteByUserID(userID uuid.UUID) error

	// Exists checks if a comment with the given ID exists
	Exists(id uuid.UUID) (bool, error)

	// FindMultiple retrieves multiple comments by their IDs
	FindMultiple(ids []uuid.UUID) ([]*entities.Comment, error)

	// Count returns the total number of comments
	Count() (int, error)

	// CountByPostID returns the number of comments for a post
	CountByPostID(postID uuid.UUID) (int, error)

	// CountByUserID returns the number of comments made by a user
	CountByUserID(userID uuid.UUID) (int, error)

	// CountByDateRange returns the number of comments within a date range
	CountByDateRange(start, end time.Time) (int, error)

	// FindRecentByPost retrieves recent comments for a post
	FindRecentByPost(postID uuid.UUID, limit int) ([]*entities.Comment, error)

	// FindRecentByUser retrieves recent comments made by a user
	FindRecentByUser(userID uuid.UUID, limit int) ([]*entities.Comment, error)

	// FindByDateRange retrieves comments within a date range
	FindByDateRange(start, end time.Time, limit, offset int) ([]*entities.Comment, error)

	// Search performs a text search across comment content
	Search(query string, limit, offset int) ([]*entities.Comment, error)

	// SearchInPost performs a text search within comments for a specific post
	SearchInPost(postID uuid.UUID, query string, limit, offset int) ([]*entities.Comment, error)

	// FindByUserIDs retrieves comments made by any of the given users
	FindByUserIDs(userIDs []uuid.UUID, limit, offset int) ([]*entities.Comment, error)

	// GetCommentersByPost returns the user IDs of users who commented on a post
	GetCommentersByPost(postID uuid.UUID, limit, offset int) ([]uuid.UUID, error)

	// GetMostCommentedPosts returns post IDs with the most comments
	GetMostCommentedPosts(limit, offset int) ([]uuid.UUID, error)
}