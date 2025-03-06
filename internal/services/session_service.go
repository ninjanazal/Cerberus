package services

import (
	"cerberus/internal/database"
	"cerberus/internal/dto/session_dto"
	"cerberus/internal/models"
	"cerberus/internal/repository"
	"cerberus/internal/tools/logger"
)

// LoginUser generates and stores JWT and refresh tokens for a user.
//
// This function performs the following steps:
// 1. Generates a JWT token for the user.
// 2. Generates a refresh token for the user.
// 3. Stores the JWT token in Redis.
// 4. Stores the refresh token in Redis.
//
// If any step fails, the function will log the error and return nil with the error.
// If the refresh token storage fails, it will also revoke the previously stored JWT token.
//
// Parameters:
//   - p_db: A pointer to the DataRefs structure containing database and configuration references.
//   - p_usr: A pointer to the User model representing the user to be logged in.
//
// Returns:
//   - *session_dto.LoginData: A pointer to LoginData containing the generated tokens if successful.
//   - error: An error if any step in the login process fails, nil otherwise.
func LoginUser(p_db *database.DataRefs, p_usr *models.User) (*session_dto.LoginData, error) {
	tkn, err := p_db.JWTGen.GenerateJWT(p_usr.ID.String())
	if err != nil {
		logger.Log("Failed to generate the JWT token - "+err.Error(), logger.ERROR)
		return nil, err
	}

	rTkn, err := p_db.JWTGen.GenerateRefreshToken()
	if err != nil {
		logger.Log("Failed to generate the JWT Refresh token - "+err.Error(), logger.ERROR)
		return nil, err
	}

	err = repository.StoreJWTToken(p_db.Redis, p_usr.ID.String(), tkn, p_db.ConfigData.RedisData.GetJWTDuration())
	if err != nil {
		logger.Log("Failed to store JWTToken - "+err.Error(), logger.ERROR)
		return nil, err
	}

	err = repository.StoreRefreshToken(p_db.Redis, p_usr.ID.String(), rTkn,
		p_db.ConfigData.RedisData.GetRefreshJWTDuration())
	if err != nil {
		repository.RevokeJWTToken(p_db.Redis, p_usr.ID.String())
		logger.Log("Failed to store RefreshToken - "+err.Error(), logger.ERROR)
		return nil, err
	}

	return &session_dto.LoginData{
		AccessToken:  tkn,
		RefreshToken: rTkn,
	}, nil
}

// RevokeAllSessionTokensToUser revokes both the JWT and refresh tokens for a given user.
//
// This function is typically used when logging out a user or invalidating their session.
// It attempts to revoke both tokens, even if one revocation fails.
//
// Parameters:
//   - p_db: A pointer to the RedisPack structure for database operations.
//   - p_usr: A pointer to the User model representing the user whose tokens should be revoked.
//
// Note: This function does not return any error. Failures in token revocation are handled silently.
// TODO: This should also return an error
func RevokeAllSessionTokensToUser(p_db *database.RedisPack, p_usr *models.User) {
	repository.RevokeJWTToken(p_db, p_usr.ID.String())
	repository.RevokeRefreshToken(p_db, p_usr.ID.String())
}

// IsTokenActive checks if a given JWT token is active for a specific user.
// It fetches the stored token for the user from Redis and compares it with the provided token.
// If the tokens match, the token is considered active.
//
// Parameters:
//   - p_db: A pointer to the database references (*database.DataRefs) containing the Redis connection.
//   - p_usrId: The unique ID (string) of the user whose token is being checked.
//   - p_tkn: The JWT token (string) to be validated.
//
// Returns:
//   - bool: `true` if the token is active (matches the stored token); otherwise, `false`.
//   - error: An error object if the token fetch operation fails; otherwise, nil.
func IsTokenActive(p_db *database.DataRefs, p_usrId string, p_tkn string) (bool, error) {
	tkn, err := repository.GetJWTToken(p_db.Redis, p_usrId)
	if err != nil {
		logger.Log("Failed to fetch the JWT token - "+err.Error(), logger.ERROR)
		return false, err
	}

	return p_tkn == tkn, nil
}
