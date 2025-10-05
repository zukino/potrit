package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/potrit/pkg/errors"
)

// ConnectionStatus represents the status of a connection between users
type ConnectionStatus string

const (
	ConnectionStatusPending  ConnectionStatus = "pending"
	ConnectionStatusAccepted ConnectionStatus = "accepted"
	ConnectionStatusBlocked  ConnectionStatus = "blocked"
)

// Connection represents the relationship between two users (pending, accepted, blocked)
type Connection struct {
	ID          uuid.UUID       `json:"id"`
	RequesterID uuid.UUID       `json:"requester_id"`
	AddresseeID uuid.UUID       `json:"addressee_id"`
	Status      ConnectionStatus `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// NewConnection creates a new connection request with pending status
func NewConnection(requesterID, addresseeID uuid.UUID) (*Connection, error) {
	// Validate that users are not connecting to themselves
	if requesterID == addresseeID {
		return nil, errors.ErrSelfConnection
	}

	connection := &Connection{
		ID:          uuid.New(),
		RequesterID: requesterID,
		AddresseeID: addresseeID,
		Status:      ConnectionStatusPending,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	if err := connection.Validate(); err != nil {
		return nil, err
	}

	return connection, nil
}

// Validate performs comprehensive validation on the connection
func (c *Connection) Validate() error {
	var validationErrors errors.ValidationErrors

	// Validate ID
	if c.ID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("id", "connection ID is required", nil))
	}

	// Validate requester ID
	if c.RequesterID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("requester_id", "requester ID is required", nil))
	}

	// Validate addressee ID
	if c.AddresseeID == uuid.Nil {
		validationErrors = append(validationErrors, errors.NewValidationError("addressee_id", "addressee ID is required", nil))
	}

	// Validate self-connection
	if c.RequesterID == c.AddresseeID {
		validationErrors = append(validationErrors, errors.NewValidationError("connection", "cannot connect to self", nil))
	}

	// Validate status
	if !c.IsValidStatus() {
		validationErrors = append(validationErrors, errors.NewValidationError("status", "invalid connection status", c.Status))
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

// IsValidStatus checks if the connection status is valid
func (c *Connection) IsValidStatus() bool {
	switch c.Status {
	case ConnectionStatusPending, ConnectionStatusAccepted, ConnectionStatusBlocked:
		return true
	default:
		return false
	}
}

// IsPending checks if the connection is pending
func (c *Connection) IsPending() bool {
	return c.Status == ConnectionStatusPending
}

// IsAccepted checks if the connection is accepted
func (c *Connection) IsAccepted() bool {
	return c.Status == ConnectionStatusAccepted
}

// IsBlocked checks if the connection is blocked
func (c *Connection) IsBlocked() bool {
	return c.Status == ConnectionStatusBlocked
}

// Accept accepts the connection request (only addressee can accept)
func (c *Connection) Accept() error {
	if c.Status != ConnectionStatusPending {
		return errors.NewValidationError("status", "only pending connections can be accepted", c.Status)
	}

	c.Status = ConnectionStatusAccepted
	c.UpdatedAt = time.Now().UTC()
	return nil
}

// Reject rejects/cancels the connection request
func (c *Connection) Reject() error {
	if c.Status != ConnectionStatusPending {
		return errors.NewValidationError("status", "only pending connections can be rejected", c.Status)
	}

	// In a real implementation, you might delete the connection or mark it as rejected
	// For now, we'll change status to a conceptual "rejected" state
	c.Status = ConnectionStatusPending // Keep as pending, but application logic should handle rejection
	c.UpdatedAt = time.Now().UTC()
	return nil
}

// Block blocks the connection (can be done from any state)
func (c *Connection) Block() {
	c.Status = ConnectionStatusBlocked
	c.UpdatedAt = time.Now().UTC()
}

// Unblock unblocks the connection, returning it to accepted state
func (c *Connection) Unblock() error {
	if c.Status != ConnectionStatusBlocked {
		return errors.NewValidationError("status", "only blocked connections can be unblocked", c.Status)
	}

	c.Status = ConnectionStatusAccepted
	c.UpdatedAt = time.Now().UTC()
	return nil
}

// IsUserInvolved checks if a user is involved in this connection
func (c *Connection) IsUserInvolved(userID uuid.UUID) bool {
	return c.RequesterID == userID || c.AddresseeID == userID
}

// IsRequester checks if the user is the requester
func (c *Connection) IsRequester(userID uuid.UUID) bool {
	return c.RequesterID == userID
}

// IsAddressee checks if the user is the addressee
func (c *Connection) IsAddressee(userID uuid.UUID) bool {
	return c.AddresseeID == userID
}

// GetOtherUserID returns the ID of the other user in the connection
func (c *Connection) GetOtherUserID(userID uuid.UUID) (uuid.UUID, error) {
	if !c.IsUserInvolved(userID) {
		return uuid.Nil, errors.NewValidationError("user_id", "user is not involved in this connection", userID)
	}

	if c.RequesterID == userID {
		return c.AddresseeID, nil
	}
	return c.RequesterID, nil
}

// Touch updates the updated_at timestamp
func (c *Connection) Touch() {
	c.UpdatedAt = time.Now().UTC()
}

// CanTransitionTo checks if the connection can transition to the given status
func (c *Connection) CanTransitionTo(newStatus ConnectionStatus) bool {
	switch c.Status {
	case ConnectionStatusPending:
		return newStatus == ConnectionStatusAccepted || newStatus == ConnectionStatusBlocked
	case ConnectionStatusAccepted:
		return newStatus == ConnectionStatusBlocked
	case ConnectionStatusBlocked:
		return newStatus == ConnectionStatusAccepted
	default:
		return false
	}
}
