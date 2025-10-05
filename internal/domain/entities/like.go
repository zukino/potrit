package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/pkg/errors"
)

// Like represents a user's positive reaction to a post
type Like struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

// NewLike creates a new like for a post by a user
func NewLike(userID, postID uuid.UUID) (*Like, error) {
	// Validate that user is not liking their own post
	if userID == postID {
		// Note: This is a basic check. In practice, we'd need to check the post's author
		return nil, errors.ErrSelfLike
	}

	like := &Like{
		ID:        uuid.New(),
		UserID:    userID,
		PostID:    postID,
		CreatedAt: time.Now().UTC(),
	}

	if err := like.Validate(); err != nil {
		return nil, err
	}

	return like, nil
}

// Validate performs comprehensive validation on the like
func (l *Like) Validate() error {
	var validationErrors errors.ValidationErrors

	// Validate ID
	if l.ID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("id", "like ID is required", nil))
	}

	// Validate user ID
	if l.UserID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("user_id", "user ID is required", nil))
	}

	// Validate post ID
	if l.PostID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("post_id", "post ID is required", nil))
	}

	// Validate created_at timestamp
	if l.CreatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("created_at", "created_at is required", nil))
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

// IsOwner checks if the given user is the owner of this like
func (l *Like) IsOwner(userID uuid.UUID) bool {
	return l.UserID == userID
}

// CanBeDeletedBy checks if the like can be deleted by the given user
func (l *Like) CanBeDeletedBy(userID uuid.UUID) bool {
	return l.IsOwner(userID)
}

// IsForPost checks if this like is for the given post
func (l *Like) IsForPost(postID uuid.UUID) bool {
	return l.PostID == postID
}

// WasCreatedBy checks if this like was created by the given user
func (l *Like) WasCreatedBy(userID uuid.UUID) bool {
	return l.UserID == userID
}

// GetAge returns the age of the like (time since creation)
func (l *Like) GetAge() time.Duration {
	return time.Since(l.CreatedAt)
}

// IsRecent checks if the like was created within the given duration
func (l *Like) IsRecent(duration time.Duration) bool {
	return time.Since(l.CreatedAt) <= duration
}
