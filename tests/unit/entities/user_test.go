package entities

import (
	"testing"

	"github.com/potrit/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// TODO: Implement User entity tests
	user := &entities.User{}
	assert.NotNil(t, user)
}

func TestUserValidation(t *testing.T) {
	// TODO: Implement user validation tests
	user := &entities.User{}
	assert.NotNil(t, user, "User should not be nil")
}