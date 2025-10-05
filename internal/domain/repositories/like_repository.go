package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/internal/domain/entities"
)

// LikeRepository defines the interface for like data operations
type LikeRepository interface {
	// Create saves a new like to the repository
	Create(like *entities.Like) error

	// FindByID retrieves a like by their ID
	FindByID(id uuid.UUID) (*entities.Like, error)

	// FindByUserAndPost retrieves a like by user and post
	FindByUserAndPost(userID, postID uuid.UUID) (*entities.Like, error)

	// FindByUserID retrieves all likes made by a user
	FindByUserID(userID uuid.UUID, limit, offset int) ([]*entities.Like, error)

	// FindByPostID retrieves all likes for a post
	FindByPostID(postID uuid.UUID, limit, offset int) ([]*entities.Like, error)

	// Delete removes a like from the repository
	Delete(id uuid.UUID) error

	// DeleteByUserAndPost removes a specific like by user and post
	DeleteByUserAndPost(userID, postID uuid.UUID) error

	// Exists checks if a like with the given ID exists
	Exists(id uuid.UUID) (bool, error)

	// ExistsByUserAndPost checks if a user has liked a specific post
	ExistsByUserAndPost(userID, postID uuid.UUID) (bool, error)

	// Count returns the total number of likes
	Count() (int, error)

	// CountByPostID returns the number of likes for a post
	CountByPostID(postID uuid.UUID) (int, error)

	// CountByUserID returns the number of likes made by a user
	CountByUserID(userID uuid.UUID) (int, error)

	// CountByDateRange returns the number of likes within a date range
	CountByDateRange(start, end time.Time) (int, error)

	// FindRecentByPost retrieves recent likes for a post
	FindRecentByPost(postID uuid.UUID, limit int) ([]*entities.Like, error)

	// FindRecentByUser retrieves recent likes made by a user
	FindRecentByUser(userID uuid.UUID, limit int) ([]*entities.Like, error)

	// FindByDateRange retrieves likes within a date range
	FindByDateRange(start, end time.Time, limit, offset int) ([]*entities.Like, error)

	// GetUsersWhoLiked returns the user IDs of users who liked a post
	GetUsersWhoLiked(postID uuid.UUID, limit, offset int) ([]uuid.UUID, error)

	// GetPostsLikedByUser returns the post IDs liked by a user
	GetPostsLikedByUser(userID uuid.UUID, limit, offset int) ([]uuid.UUID, error)
}