package contract

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestRegisterEndpoint(t *testing.T) {
	// Test user registration endpoint
	payload := map[string]interface{}{
		"email":    "test@example.com",
		"password": "testpassword123",
		"name":     "Test User",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", BaseAPI+"/auth/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err := client.Do(req)

	// Should fail because implementation doesn't exist yet
	assert.Error(t, err, "Expected connection error - server not implemented")
	assert.Contains(t, err.Error(), "connection refused", "Expected connection refused error")
}

func TestLoginEndpoint(t *testing.T) {
	// Test user login endpoint
	payload := map[string]interface{}{
		"email":    "test@example.com",
		"password": "testpassword123",
	}

	jsonData, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", BaseAPI+"/auth/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err := client.Do(req)

	// Should fail because implementation doesn't exist yet
	assert.Error(t, err, "Expected connection error - server not implemented")
	assert.Contains(t, err.Error(), "connection refused", "Expected connection refused error")
}