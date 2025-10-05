package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// PasswordService handles password hashing and verification
type PasswordService struct {
	cost int
}

// NewPasswordService creates a new password service
func NewPasswordService() *PasswordService {
	return &PasswordService{
		cost: bcrypt.DefaultCost,
	}
}

// NewPasswordServiceWithCost creates a new password service with custom cost
func NewPasswordServiceWithCost(cost int) *PasswordService {
	if cost < bcrypt.MinCost || cost > bcrypt.MaxCost {
		cost = bcrypt.DefaultCost
	}
	return &PasswordService{
		cost: cost,
	}
}

// HashPassword hashes a password using bcrypt
func (s *PasswordService) HashPassword(password string) (string, error) {
	if err := s.validatePassword(password); err != nil {
		return "", fmt.Errorf("password validation failed: %w", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), s.cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hash), nil
}

// CheckPassword verifies a password against a hash
func (s *PasswordService) CheckPassword(password, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return fmt.Errorf("password verification failed: %w", err)
	}
	return nil
}

// CheckPasswordSecure securely verifies a password against a hash with timing attack protection
func (s *PasswordService) CheckPasswordSecure(password, hash string) error {
	// Hash the provided password to compare lengths
	if err := s.validatePassword(password); err != nil {
		return fmt.Errorf("password validation failed: %w", err)
	}

	// Use bcrypt's constant-time comparison
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return fmt.Errorf("password verification failed: %w", err)
	}

	return nil
}

// GenerateRandomPassword generates a secure random password
func (s *PasswordService) GenerateRandomPassword(length int) (string, error) {
	if length < 8 {
		length = 8
	}
	if length > 128 {
		length = 128
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random password: %w", err)
	}

	for i, b := range bytes {
		bytes[i] = charset[b%byte(len(charset))]
	}

	return string(bytes), nil
}

// ValidatePassword checks if a password meets security requirements
func (s *PasswordService) ValidatePassword(password string) error {
	return s.validatePassword(password)
}

// validatePassword performs the actual password validation
func (s *PasswordService) validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if len(password) > 128 {
		return fmt.Errorf("password must not exceed 128 characters")
	}

	// Check for at least one lowercase letter
	if matched, _ := regexp.MatchString("[a-z]", password); !matched {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	// Check for at least one uppercase letter
	if matched, _ := regexp.MatchString("[A-Z]", password); !matched {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	// Check for at least one digit
	if matched, _ := regexp.MatchString("[0-9]", password); !matched {
		return fmt.Errorf("password must contain at least one digit")
	}

	// Check for at least one special character
	if matched, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, password); !matched {
		return fmt.Errorf("password must contain at least one special character")
	}

	// Check for common weak passwords (basic check)
	weakPasswords := []string{
		"password", "12345678", "qwerty123", "admin123", "password123",
		"123456789", "password1", "welcome123", "password123!",
	}

	lowercasePassword := toLower(password)
	for _, weak := range weakPasswords {
		if subtle.ConstantTimeCompare([]byte(lowercasePassword), []byte(weak)) == 1 {
			return fmt.Errorf("password is too common and weak")
		}
	}

	return nil
}

// GetStrength returns a password strength score (0-100)
func (s *PasswordService) GetStrength(password string) int {
	score := 0

	// Length score
	if len(password) >= 8 {
		score += 20
	}
	if len(password) >= 12 {
		score += 10
	}
	if len(password) >= 16 {
		score += 10
	}

	// Character variety score
	if matched, _ := regexp.MatchString("[a-z]", password); matched {
		score += 10
	}
	if matched, _ := regexp.MatchString("[A-Z]", password); matched {
		score += 10
	}
	if matched, _ := regexp.MatchString("[0-9]", password); matched {
		score += 10
	}
	if matched, _ := regexp.MatchString(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`, password); matched {
		score += 15
	}

	// Complexity bonus
	if matched, _ := regexp.MatchString(`^.{8,}$`, password); matched {
		score += 5
	}
	if matched, _ := regexp.MatchString(`.{12,}$`, password); matched {
		score += 5
	}
	if matched, _ := regexp.MatchString(`.{16,}$`, password); matched {
		score += 5
	}

	if score > 100 {
		score = 100
	}

	return score
}

// GetStrengthDescription returns a human-readable strength description
func (s *PasswordService) GetStrengthDescription(password string) string {
	strength := s.GetStrength(password)
	switch {
	case strength >= 80:
		return "Very Strong"
	case strength >= 60:
		return "Strong"
	case strength >= 40:
		return "Medium"
	case strength >= 20:
		return "Weak"
	default:
		return "Very Weak"
	}
}

// helper function to convert string to lowercase
func toLower(s string) string {
	result := make([]byte, len(s))
	for i, b := range []byte(s) {
		if b >= 'A' && b <= 'Z' {
			result[i] = b + 32
		} else {
			result[i] = b
		}
	}
	return string(result)
}