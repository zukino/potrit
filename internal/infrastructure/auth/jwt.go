package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// TokenService handles JWT token generation and validation
type TokenService struct {
	secret         []byte
	issuer         string
	expiration     time.Duration
	refreshTime    time.Duration
}

// Claims represents the JWT claims structure
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// TokenPair represents an access token and refresh token pair
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// NewTokenService creates a new JWT token service
func NewTokenService(secret string, expiration int, issuer string) (*TokenService, error) {
	if len(secret) < 32 {
		return nil, fmt.Errorf("JWT secret must be at least 32 characters")
	}

	if expiration <= 0 {
		expiration = 60 // default 1 hour
	}

	if issuer == "" {
		issuer = "social-media-api"
	}

	return &TokenService{
		secret:      []byte(secret),
		issuer:      issuer,
		expiration:  time.Duration(expiration) * time.Minute,
		refreshTime: time.Duration(expiration*24) * time.Minute, // 24x expiration for refresh
	}, nil
}

// GenerateTokenPair generates an access token and refresh token for a user
func (s *TokenService) GenerateTokenPair(userID, email, name string) (*TokenPair, error) {
	if userID == "" || email == "" {
		return nil, fmt.Errorf("user ID and email are required")
	}

	now := time.Now()
	expiresAt := now.Add(s.expiration)

	// Create claims
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Issuer:    s.issuer,
			Subject:   userID,
			Audience:  []string{"social-media-app"},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	// Create access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(s.secret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Create refresh token
	refreshClaims := &Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			Issuer:    s.issuer,
			Subject:   userID,
			Audience:  []string{"social-media-app"},
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTime)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSigned, err := refreshToken.SignedString(s.secret)
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshSigned,
		ExpiresAt:    expiresAt,
	}, nil
}

// ValidateToken validates an access token and returns the claims
func (s *TokenService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Check if token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ValidateRefreshToken validates a refresh token and returns the claims
func (s *TokenService) ValidateRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse refresh token: %w", err)
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// Check if token is expired
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("refresh token has expired")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("invalid refresh token")
}

// RefreshAccessToken generates a new access token from a refresh token
func (s *TokenService) RefreshAccessToken(refreshTokenString string) (*TokenPair, error) {
	claims, err := s.ValidateRefreshToken(refreshTokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	return s.GenerateTokenPair(claims.UserID, claims.Email, claims.Name)
}

// ExtractTokenFromHeader extracts token from Authorization header
func (s *TokenService) ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is empty")
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return authHeader[len(bearerPrefix):], nil
}

// GenerateSecureSecret generates a cryptographically secure random secret
func GenerateSecureSecret(length int) (string, error) {
	if length < 32 {
		length = 32
	}

	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate secure secret: %w", err)
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GetTokenExpiration returns the expiration time of the token
func (s *TokenService) GetTokenExpiration() time.Duration {
	return s.expiration
}

// GetRefreshExpiration returns the refresh token expiration time
func (s *TokenService) GetRefreshExpiration() time.Duration {
	return s.refreshTime
}

// IsTokenExpired checks if a token is expired
func IsTokenExpired(claims *Claims) bool {
	if claims == nil || claims.ExpiresAt == nil {
		return true
	}
	return claims.ExpiresAt.Time.Before(time.Now())
}

// GetTokenRemainingTime returns the remaining time until token expires
func GetTokenRemainingTime(claims *Claims) time.Duration {
	if claims == nil || claims.ExpiresAt == nil {
		return 0
	}
	return time.Until(claims.ExpiresAt.Time)
}