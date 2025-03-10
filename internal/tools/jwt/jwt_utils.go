package jwt

import (
	"cerberus/internal/tools/logger"
	"cerberus/pkg/config"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims represents the custom claims structure for JWT tokens.
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTGenerator is responsible for generating and validating JWT tokens.
type JWTGenerator struct {
	Secret   []byte
	Duration time.Duration
}

// NewJWTGenerator creates a new JWTGenerator instance with the provided configuration.
// It sets the JWT secret from the environment variable and parses the duration from the config.
// If parsing fails, it defaults to 15 minutes.
//
// Parameters:
//   - p_cfg: A pointer to the ConfigData structure containing Redis configuration.
//
// Returns:
//   - *JWTGenerator: A pointer to the newly created JWTGenerator instance.
func NewJWTGenerator(p_cfg *config.ConfigData) *JWTGenerator {
	d, err := time.ParseDuration(p_cfg.RedisData.JWTDuration)
	if err != nil {
		logger.Log("Failed to parse JWT configuration, fail to default", logger.ERROR)
		d = 15
	}
	return &JWTGenerator{
		Secret:   []byte(os.Getenv("JWT_SECRET")),
		Duration: d,
	}
}

// GenerateJWT generates a new JWT token for the given user ID.
//
// Parameters:
//   - p_usrId: The user ID to be included in the token claims.
//
// Returns:
//   - string: The generated JWT token as a string.
//   - error: An error if token generation fails, nil otherwise.
func (gen *JWTGenerator) GenerateJWT(p_usrId string) (string, error) {
	var expiration time.Time = time.Now().Add(gen.Duration)

	claims := &Claims{
		UserID: p_usrId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}

	var tkn *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tkn.SignedString(gen.Secret)
}

// ValidateJWT validates the provided JWT token and returns the claims if valid.
//
// Parameters:
//   - p_token: The JWT token string to validate.
//
// Returns:
//   - *Claims: A pointer to the Claims structure if the token is valid.
//   - error: An error if the token is invalid or validation fails, nil otherwise.
func (gen *JWTGenerator) ValidateJWT(p_token string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(p_token, claims, func(t *jwt.Token) (interface{}, error) {
		return gen.Secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GenerateRefreshToken generates a new refresh token.
//
// Returns:
//   - string: The generated refresh token as a base64-encoded string.
//   - error: An error if token generation fails, nil otherwise.
func (gen *JWTGenerator) GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

// ValidateRefreshToken compares the saved refresh token with the target token.
//
// Parameters:
//   - p_savedToken: The saved refresh token to compare against.
//   - p_targetToken: The target refresh token to validate.
//
// Returns:
//   - bool: True if the tokens match, false otherwise.
func (gen *JWTGenerator) ValidateRefreshToken(p_savedToken string, p_targetToken string) bool {
	return p_savedToken == p_targetToken
}

// GetUserIDFromToken extracts the user ID from the provided JWT token.
// It parses the token using the JWT secret stored in the JWTGenerator instance
// and validates the token's signature. If the token is valid, it returns the user ID
// embedded in the token. If the token is invalid or parsing fails, it returns an error.
func (gen *JWTGenerator) GetUserIDFromToken(p_token string) (string, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(p_token, claims, func(t *jwt.Token) (interface{}, error) {
		return gen.Secret, nil
	})
	if err != nil {
		return "", errors.New("invalid token")
	}

	return claims.UserID, nil
}
