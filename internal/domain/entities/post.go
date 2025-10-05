package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/pkg/errors"
)

// Post represents content created by users containing text content and interaction metadata
// (media support excluded from MVP)
type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost creates a new post with the given user ID and content
func NewPost(userID uuid.UUID, content string) (*Post, error) {
	post := &Post{
		ID:        uuid.New(),
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := post.Validate(); err != nil {
		return nil, err
	}

	return post, nil
}

// Validate performs comprehensive validation on the post
func (p *Post) Validate() error {
	var validationErrors errors.ValidationErrors

	// Validate ID
	if p.ID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("id", "post ID is required", nil))
	}

	// Validate user ID
	if p.UserID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("user_id", "user ID is required", nil))
	}

	// Validate content
	if err := errors.ValidateContent(p.Content, "content", 1, 2000); err != nil {
		validationErrors = append(validationErrors, err.(errors.ValidationError))
	}

	// Validate timestamps
	if p.CreatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("created_at", "created_at is required", nil))
	}

	if p.UpdatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("updated_at", "updated_at is required", nil))
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

// UpdateContent updates the post content with validation
func (p *Post) UpdateContent(content string) error {
	if err := errors.ValidateContent(content, "content", 1, 2000); err != nil {
		return err
	}

	p.Content = content
	p.UpdatedAt = time.Now().UTC()
	return nil
}

// IsAuthor checks if the given user ID is the author of this post
func (p *Post) IsAuthor(userID uuid.UUID) bool {
	return p.UserID == userID
}

// CanBeEditedBy checks if the post can be edited by the given user
func (p *Post) CanBeEditedBy(userID uuid.UUID) bool {
	return p.IsAuthor(userID)
}

// Touch updates the updated_at timestamp
func (p *Post) Touch() {
	p.UpdatedAt = time.Now().UTC()
}

// GetContentLength returns the length of the post content
func (p *Post) GetContentLength() int {
	return len(p.Content)
}

// IsEmpty checks if the post content is empty
func (p *Post) IsEmpty() bool {
	return len(p.Content) == 0
}
