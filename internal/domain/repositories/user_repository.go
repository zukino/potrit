package repositories

import (
	"github.com/google/uuid"
	"github.com/potrit/internal/domain/entities"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	// Create saves a new user to the repository
	Create(user *entities.User) error

	// FindByID retrieves a user by their ID
	FindByID(id uuid.UUID) (*entities.User, error)

	// FindByEmail retrieves a user by their email address
	FindByEmail(email string) (*entities.User, error)

	// FindByName retrieves users by name (partial match)
	FindByName(name string, limit, offset int) ([]*entities.User, error)

	// Search performs a text search across user names and emails
	Search(query string, limit, offset int) ([]*entities.User, error)

	// Update updates an existing user in the repository
	Update(user *entities.User) error

	// Delete removes a user from the repository
	Delete(id uuid.UUID) error

	// Exists checks if a user with the given ID exists
	Exists(id uuid.UUID) (bool, error)

	// ExistsByEmail checks if a user with the given email exists
	ExistsByEmail(email string) (bool, error)

	// FindMultiple retrieves multiple users by their IDs
	FindMultiple(ids []uuid.UUID) ([]*entities.User, error)

	// Count returns the total number of users
	Count() (int, error)

	// CountByStatus returns the number of users by status
	CountByStatus(status entities.UserStatus) (int, error)

	// FindActive returns only active users
	FindActive(limit, offset int) ([]*entities.User, error)

	// FindByConnectionStatus finds users based on their connection status with a given user
	FindByConnectionStatus(userID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.User, error)
}
