package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/internal/domain/entities"
)

// PostRepository defines the interface for post data operations
type PostRepository interface {
	// Create saves a new post to the repository
	Create(post *entities.Post) error

	// FindByID retrieves a post by their ID
	FindByID(id uuid.UUID) (*entities.Post, error)

	// FindByUserID retrieves posts created by a specific user
	FindByUserID(userID uuid.UUID, limit, offset int) ([]*entities.Post, error)

	// FindByUserIDs retrieves posts created by any of the given users
	FindByUserIDs(userIDs []uuid.UUID, limit, offset int) ([]*entities.Post, error)

	// FindFeed retrieves posts for a user's feed (from their connections)
	FindFeed(userID uuid.UUID, limit, offset int) ([]*entities.Post, error)

	// FindByDateRange retrieves posts within a date range
	FindByDateRange(start, end time.Time, limit, offset int) ([]*entities.Post, error)

	// Search performs a text search across post content
	Search(query string, limit, offset int) ([]*entities.Post, error)

	// Update updates an existing post in the repository
	Update(post *entities.Post) error

	// Delete removes a post from the repository
	Delete(id uuid.UUID) error

	// Exists checks if a post with the given ID exists
	Exists(id uuid.UUID) (bool, error)

	// FindMultiple retrieves multiple posts by their IDs
	FindMultiple(ids []uuid.UUID) ([]*entities.Post, error)

	// Count returns the total number of posts
	Count() (int, error)

	// CountByUserID returns the number of posts by a specific user
	CountByUserID(userID uuid.UUID) (int, error)

	// FindRecent retrieves the most recent posts
	FindRecent(limit, offset int) ([]*entities.Post, error)

	// FindPopular retrieves posts with the most interactions (likes/comments)
	FindPopular(limit, offset int) ([]*entities.Post, error)

	// GetLikeCount returns the number of likes for a post
	GetLikeCount(postID uuid.UUID) (int, error)

	// GetCommentCount returns the number of comments for a post
	GetCommentCount(postID uuid.UUID) (int, error)
}
