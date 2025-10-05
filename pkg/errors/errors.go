package errors

import (
	"errors"
	"fmt"
)

// Domain errors for the social media application
var (
	// User related errors
	ErrUserNotFound         = errors.New("user not found")
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUserEmailInvalid     = errors.New("invalid email format")
	ErrUserNameInvalid      = errors.New("invalid name format")
	ErrUserPasswordInvalid  = errors.New("invalid password format")
	ErrUserSelfOperation    = errors.New("cannot perform operation on self")

	// Post related errors
	ErrPostNotFound       = errors.New("post not found")
	ErrPostAccessDenied   = errors.New("access denied to post")
	ErrPostContentInvalid = errors.New("invalid post content")
	ErrPostTooLong        = errors.New("post content exceeds maximum length")

	// Connection related errors
	ErrConnectionNotFound      = errors.New("connection not found")
	ErrConnectionAlreadyExists = errors.New("connection already exists")
	ErrConnectionLimitReached  = errors.New("connection limit reached")
	ErrInvalidConnectionStatus = errors.New("invalid connection status")
	ErrSelfConnection          = errors.New("cannot connect to self")

	// Like related errors
	ErrLikeNotFound        = errors.New("like not found")
	ErrLikeAlreadyExists   = errors.New("like already exists")
	ErrSelfLike            = errors.New("cannot like own post")

	// Comment related errors
	ErrCommentNotFound       = errors.New("comment not found")
	ErrCommentAccessDenied   = errors.New("access denied to comment")
	ErrCommentContentInvalid = errors.New("invalid comment content")
	ErrCommentTooLong        = errors.New("comment content exceeds maximum length")

	// Validation errors
	ErrInvalidInput         = errors.New("invalid input data")
	ErrRequiredFieldMissing = errors.New("required field missing")
	ErrInvalidUUID          = errors.New("invalid UUID format")
)

// ValidationError represents a validation error with field details
type ValidationError struct {
	Field   string
	Message string
	Value   any
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error on field '%s': %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	if len(ve) == 0 {
		return "no validation errors"
	}
	if len(ve) == 1 {
		return ve[0].Error()
	}

	msg := fmt.Sprintf("multiple validation errors (%d):", len(ve))
	for _, err := range ve {
		msg += "\n  - " + err.Error()
	}
	return msg
}

// NewValidationError creates a new validation error
func NewValidationError(field, message string, value any) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	}
}

// Validation helpers

// ValidateEmail checks if email format is valid
func ValidateEmail(email string) error {
	// Basic email validation - can be enhanced with regex
	if len(email) < 3 || len(email) > 254 {
		return NewValidationError("email", "email must be between 3 and 254 characters", email)
	}

	// Check for @ symbol and basic structure
	atIndex := -1
	for i, char := range email {
		if char == '@' {
			if atIndex != -1 {
				return NewValidationError("email", "email cannot contain multiple @ symbols", email)
			}
			atIndex = i
		}
	}

	if atIndex == -1 || atIndex == 0 || atIndex == len(email)-1 {
		return NewValidationError("email", "email must contain @ symbol with valid local and domain parts", email)
	}

	return nil
}

// ValidatePassword checks if password meets requirements
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return NewValidationError("password", "password must be at least 8 characters long", len(password))
	}

	if len(password) > 128 {
		return NewValidationError("password", "password must not exceed 128 characters", len(password))
	}

	return nil
}

// ValidateName checks if name meets requirements
func ValidateName(name string) error {
	if len(name) == 0 {
		return NewValidationError("name", "name is required", name)
	}

	if len(name) > 100 {
		return NewValidationError("name", "name must not exceed 100 characters", len(name))
	}

	return nil
}

// ValidateContent checks if content meets requirements
func ValidateContent(content string, fieldName string, minLength, maxLength int) error {
	if len(content) < minLength {
		return NewValidationError(fieldName, fmt.Sprintf("%s must be at least %d characters", fieldName, minLength), len(content))
	}

	if len(content) > maxLength {
		return NewValidationError(fieldName, fmt.Sprintf("%s must not exceed %d characters", fieldName, maxLength), len(content))
	}

	return nil
}

// ValidateUUID checks if UUID format is valid (basic check)
func ValidateUUID(uuid string) error {
	if len(uuid) != 36 {
		return NewValidationError("id", "invalid UUID format", uuid)
	}

	// Basic UUID format validation - can be enhanced
	hyphenCount := 0
	for _, char := range uuid {
		if char == '-' {
			hyphenCount++
		}
	}

	if hyphenCount != 4 {
		return NewValidationError("id", "invalid UUID format", uuid)
	}

	return nil
}
