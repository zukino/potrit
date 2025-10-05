package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/pkg/errors"
)

// Comment represents a user's textual response to a post
type Comment struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewComment creates a new comment on a post by a user
func NewComment(userID, postID uuid.UUID, content string) (*Comment, error) {
	comment := &Comment{
		ID:        uuid.New(),
		UserID:    userID,
		PostID:    postID,
		Content:   content,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := comment.Validate(); err != nil {
		return nil, err
	}

	return comment, nil
}

// Validate performs comprehensive validation on the comment
func (c *Comment) Validate() error {
	var validationErrors errors.ValidationErrors

	// Validate ID
	if c.ID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("id", "comment ID is required", nil))
	}

	// Validate user ID
	if c.UserID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("user_id", "user ID is required", nil))
	}

	// Validate post ID
	if c.PostID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("post_id", "post ID is required", nil))
	}

	// Validate content
	if err := errors.ValidateContent(c.Content, "content", 1, 1000); err != nil {
		validationErrors = append(validationErrors, err.(errors.ValidationError))
	}

	// Validate timestamps
	if c.CreatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("created_at", "created_at is required", nil))
	}

	if c.UpdatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("updated_at", "updated_at is required", nil))
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

// UpdateContent updates the comment content with validation
func (c *Comment) UpdateContent(content string) error {
	if err := errors.ValidateContent(content, "content", 1, 1000); err != nil {
		return err
	}

	c.Content = content
	c.UpdatedAt = time.Now().UTC()
	return nil
}

// IsAuthor checks if the given user is the author of this comment
func (c *Comment) IsAuthor(userID uuid.UUID) bool {
	return c.UserID == userID
}

// CanBeEditedBy checks if the comment can be edited by the given user
func (c *Comment) CanBeEditedBy(userID uuid.UUID) bool {
	return c.IsAuthor(userID)
}

// CanBeDeletedBy checks if the comment can be deleted by the given user
func (c *Comment) CanBeDeletedBy(userID uuid.UUID) bool {
	return c.IsAuthor(userID)
}

// IsForPost checks if this comment is for the given post
func (c *Comment) IsForPost(postID uuid.UUID) bool {
	return c.PostID == postID
}

// WasCreatedBy checks if this comment was created by the given user
func (c *Comment) WasCreatedBy(userID uuid.UUID) bool {
	return c.UserID == userID
}

// Touch updates the updated_at timestamp
func (c *Comment) Touch() {
	c.UpdatedAt = time.Now().UTC()
}

// GetContentLength returns the length of the comment content
func (c *Comment) GetContentLength() int {
	return len(c.Content)
}

// IsEmpty checks if the comment content is empty
func (c *Comment) IsEmpty() bool {
	return len(c.Content) == 0
}

// GetAge returns the age of the comment (time since creation)
func (c *Comment) GetAge() time.Duration {
	return time.Since(c.CreatedAt)
}

// IsRecent checks if the comment was created within the given duration
func (c *Comment) IsRecent(duration time.Duration) bool {
	return time.Since(c.CreatedAt) <= duration
}

// WasRecentlyUpdated checks if the comment was updated within the given duration
func (c *Comment) WasRecentlyUpdated(duration time.Duration) bool {
	return time.Since(c.UpdatedAt) <= duration
}

// GetWordCount returns the number of words in the comment content
func (c *Comment) GetWordCount() int {
	if len(c.Content) == 0 {
		return 0
	}

	// Simple word count - split by whitespace
	// In a production system, you might want more sophisticated word counting
	wordCount := 0
	inWord := false

	for _, char := range c.Content {
		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			if inWord {
				wordCount++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	// Count the last word if the string doesn't end with whitespace
	if inWord {
		wordCount++
	}

	return wordCount
}
