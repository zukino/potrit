package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/pkg/errors"
)

// UserStatus represents the status of a user account
type UserStatus string

const (
	UserStatusPending  UserStatus = "pending"
	UserStatusActive   UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
)

// User represents a person on the platform with profile information,
// authentication credentials, connection relationships, and equal permissions to all other users
type User struct {
	ID           uuid.UUID  `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"` // Never expose password hash in JSON
	Name         string     `json:"name"`
	Status       UserStatus `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// NewUser creates a new user with the given email, password hash, and name
func NewUser(email, passwordHash, name string) (*User, error) {
	user := &User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		Status:       UserStatusPending, // New users start as pending
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Validate performs comprehensive validation on the user
func (u *User) Validate() error {
	var validationErrors errors.ValidationErrors

	// Validate email
	if err := errors.ValidateEmail(u.Email); err != nil {
		validationErrors = append(validationErrors, err.(errors.ValidationError))
	}

	// Validate password hash
	if len(u.PasswordHash) == 0 {
		validationErrors = append(validationErrors, errors.NewValidationError("password_hash", "password hash is required", nil))
	}

	// Validate name
	if err := errors.ValidateName(u.Name); err != nil {
		validationErrors = append(validationErrors, err.(errors.ValidationError))
	}

	// Validate status
	if !u.IsValidStatus() {
		validationErrors = append(validationErrors, errors.NewValidationError("status", "invalid user status", u.Status))
	}

	// Validate timestamps
	if u.CreatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("created_at", "created_at is required", nil))
	}

	if u.UpdatedAt.IsZero() {
		validationErrors = append(validationErrors, errors.NewValidationError("updated_at", "updated_at is required", nil))
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

// IsValidStatus checks if the user status is valid
func (u *User) IsValidStatus() bool {
	switch u.Status {
	case UserStatusPending, UserStatusActive, UserStatusSuspended:
		return true
	default:
		return false
	}
}

// CanLogin checks if the user can login
func (u *User) CanLogin() bool {
	return u.Status == UserStatusActive
}

// Activate sets the user status to active
func (u *User) Activate() {
	u.Status = UserStatusActive
	u.UpdatedAt = time.Now().UTC()
}

// Suspend sets the user status to suspended
func (u *User) Suspend() {
	u.Status = UserStatusSuspended
	u.UpdatedAt = time.Now().UTC()
}

// UpdateName updates the user's name with validation
func (u *User) UpdateName(name string) error {
	if err := errors.ValidateName(name); err != nil {
		return err
	}

	u.Name = name
	u.UpdatedAt = time.Now().UTC()
	return nil
}

// UpdatePassword updates the user's password hash
func (u *User) UpdatePassword(passwordHash string) error {
	if len(passwordHash) == 0 {
		return errors.NewValidationError("password_hash", "password hash is required", nil)
	}

	u.PasswordHash = passwordHash
	u.UpdatedAt = time.Now().UTC()
	return nil
}

// UpdateEmail updates the user's email with validation
func (u *User) UpdateEmail(email string) error {
	if err := errors.ValidateEmail(email); err != nil {
		return err
	}

	u.Email = email
	u.UpdatedAt = time.Now().UTC()
	return nil
}

// IsSameAs checks if this user is the same as another user
func (u *User) IsSameAs(otherID uuid.UUID) bool {
	return u.ID == otherID
}

// Touch updates the updated_at timestamp
func (u *User) Touch() {
	u.UpdatedAt = time.Now().UTC()
}
