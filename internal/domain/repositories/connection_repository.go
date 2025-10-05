package repositories

import (
	"github.com/google/uuid"
	"github.com/potrit/internal/domain/entities"
)

// ConnectionRepository defines the interface for connection data operations
type ConnectionRepository interface {
	// Create saves a new connection to the repository
	Create(connection *entities.Connection) error

	// FindByID retrieves a connection by their ID
	FindByID(id uuid.UUID) (*entities.Connection, error)

	// FindByUsers retrieves a connection between two specific users
	FindByUsers(requesterID, addresseeID uuid.UUID) (*entities.Connection, error)

	// FindByUserID retrieves connections for a user (as requester or addressee)
	FindByUserID(userID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.Connection, error)

	// FindByRequesterID retrieves connections where the user is the requester
	FindByRequesterID(requesterID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.Connection, error)

	// FindByAddresseeID retrieves connections where the user is the addressee
	FindByAddresseeID(addresseeID uuid.UUID, status entities.ConnectionStatus, limit, offset int) ([]*entities.Connection, error)

	// FindAcceptedConnections retrieves all accepted connections for a user
	FindAcceptedConnections(userID uuid.UUID, limit, offset int) ([]*entities.Connection, error)

	// FindPendingRequests retrieves pending connection requests for a user
	FindPendingRequests(userID uuid.UUID, limit, offset int) ([]*entities.Connection, error)

	// FindSentRequests retrieves connection requests sent by a user
	FindSentRequests(userID uuid.UUID, limit, offset int) ([]*entities.Connection, error)

	// Update updates an existing connection in the repository
	Update(connection *entities.Connection) error

	// Delete removes a connection from the repository
	Delete(id uuid.UUID) error

	// Exists checks if a connection with the given ID exists
	Exists(id uuid.UUID) (bool, error)

	// ExistsBetweenUsers checks if a connection exists between two users
	ExistsBetweenUsers(requesterID, addresseeID uuid.UUID) (bool, error)

	// Count returns the total number of connections
	Count() (int, error)

	// CountByStatus returns the number of connections by status
	CountByStatus(status entities.ConnectionStatus) (int, error)

	// CountByUserID returns the number of connections for a user
	CountByUserID(userID uuid.UUID) (int, error)

	// GetConnectedUserIDs returns the IDs of users connected to the given user
	GetConnectedUserIDs(userID uuid.UUID) ([]uuid.UUID, error)

	// AreConnected checks if two users have an accepted connection
	AreConnected(userID1, userID2 uuid.UUID) (bool, error)

	// GetConnectionStatus returns the connection status between two users
	GetConnectionStatus(userID1, userID2 uuid.UUID) (entities.ConnectionStatus, error)
}
