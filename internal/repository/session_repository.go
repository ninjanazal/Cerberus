package repository

import (
	"cerberus/internal/database"
	"cerberus/internal/tools/logger"
	"errors"
	"time"
)

var (
	refreshPrefix string = "refresh:" // refreshPrefix is the prefix used for storing refresh tokens in Redis.
	tokenPrefix   string = "token:"   // tokenPrefix is the prefix used for storing JWT tokens in Redis.
)

// StoreJWTToken stores a JWT token in Redis for a given user ID.
//
// Parameters:
//   - p_db: A pointer to the RedisPack instance for database operations.
//   - p_usrId: The user ID associated with the token.
//   - p_jwtToken: The JWT token to store.
//   - p_duration: The duration for which the token should be stored.
//
// Returns:
//   - error: An error if the storage operation fails, nil otherwise.
func StoreJWTToken(p_db *database.RedisPack, p_usrId string, p_jwtToken string, p_duration time.Duration) error {
	return p_db.Client.Set(p_db.Ctx, tokenPrefix+p_usrId, p_jwtToken, p_duration).Err()
}

// GetJWTToken retrieves a JWT token from Redis for a given user ID.
//
// Parameters:
//   - p_db: A pointer to the RedisPack instance for database operations.
//   - p_usrId: The user ID associated with the token.
//
// Returns:
//   - string: The retrieved JWT token.
//   - error: An error if the token is not found or retrieval fails, nil otherwise.
func GetJWTToken(p_db *database.RedisPack, p_usrId string) (string, error) {
	tkn, err := p_db.Client.Get(p_db.Ctx, tokenPrefix+p_usrId).Result()
	if err != nil {
		logger.Log("JWT Token not found - "+err.Error(), logger.INFO)
		return "", errors.New("not found")
	}

	return tkn, nil
}

// RevokeJWTToken removes a JWT token from Redis for a given user ID.
//
// Parameters:
//   - p_db: A pointer to the RedisPack instance for database operations.
//   - p_usrId: The user ID associated with the token to be revoked.
//
// Returns:
//   - error: An error if the revocation operation fails, nil otherwise.
func RevokeJWTToken(p_db *database.RedisPack, p_usrId string) error {
	return p_db.Client.Del(p_db.Ctx, tokenPrefix+p_usrId).Err()
}

// StoreRefreshToken stores a refresh token in Redis for a given user ID.
//
// Parameters:
//   - p_db: A pointer to the RedisPack instance for database operations.
//   - p_usrId: The user ID associated with the refresh token.
//   - p_refreshToken: The refresh token to store.
//   - p_duration: The duration for which the refresh token should be stored.
//
// Returns:
//   - error: An error if the storage operation fails, nil otherwise.
func StoreRefreshToken(p_db *database.RedisPack, p_usrId string, p_refreshToken string, p_duration time.Duration) error {
	return p_db.Client.Set(p_db.Ctx, refreshPrefix+p_usrId, p_refreshToken, p_duration).Err()
}

// GetRefreshToken retrieves a refresh token from Redis for a given user ID.
//
// Parameters:
//   - p_db: A pointer to the RedisPack instance for database operations.
//   - p_usrId: The user ID associated with the refresh token.
//
// Returns:
//   - string: The retrieved refresh token.
//   - error: An error if the token is not found or retrieval fails, nil otherwise.
func GetRefreshToken(p_db *database.RedisPack, p_usrId string) (string, error) {
	tkn, err := p_db.Client.Get(p_db.Ctx, refreshPrefix+p_usrId).Result()
	if err != nil {
		logger.Log("Refresh token not found - "+err.Error(), logger.INFO)
		return "", errors.New("not found")
	}
	return tkn, nil
}

// RevokeRefreshToken removes a refresh token from Redis for a given user ID.
//
// Parameters:
//   - p_db: A pointer to the RedisPack instance for database operations.
//   - p_usrId: The user ID associated with the refresh token to be revoked.
//
// Returns:
//   - error: An error if the revocation operation fails, nil otherwise.
func RevokeRefreshToken(p_db *database.RedisPack, p_usrId string) error {
	return p_db.Client.Del(p_db.Ctx, refreshPrefix+p_usrId).Err()
}
