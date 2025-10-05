package entities

import (
	"testing"

	"github.com/potrit/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	// TODO: Implement Post entity tests
	post := &entities.Post{}
	assert.NotNil(t, post)
}

func TestPostValidation(t *testing.T) {
	// TODO: Implement post validation tests
	post := &entities.Post{}
	assert.NotNil(t, post, "Post should not be nil")
}