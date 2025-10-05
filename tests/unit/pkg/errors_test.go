package errors

import (
	"testing"

	"github.com/potrit/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	// Test that domain errors are properly defined
	assert.NotNil(t, errors.ErrUserNotFound)
	assert.NotNil(t, errors.ErrPostNotFound)
	assert.NotNil(t, errors.ErrConnectionNotFound)
	assert.NotNil(t, errors.ErrInvalidInput)
}

func TestErrorMessages(t *testing.T) {
	// Test that error messages are descriptive
	assert.Equal(t, "user not found", errors.ErrUserNotFound.Error())
	assert.Equal(t, "post not found", errors.ErrPostNotFound.Error())
	assert.Equal(t, "connection not found", errors.ErrConnectionNotFound.Error())
	assert.Equal(t, "invalid input data", errors.ErrInvalidInput.Error())
}