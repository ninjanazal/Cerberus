package session_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/dto/session_dto"
	"cerberus/internal/middleware"
	"cerberus/internal/services"
	"cerberus/internal/tools/logger"
	"encoding/json"
	"net/http"
)

// CreateRefreshHandler returns an HTTP handler function for refreshing session tokens.
// It validates the provided JWT token and refresh token, revokes all existing session tokens for the user,
// generates new tokens, and returns them in the response.
//
// Parameters:
//   - p_db: A pointer to the database.DataRefs struct containing database references.
//
// Returns:
//   - http.HandlerFunc: The HTTP handler function that processes the refresh token request.
func CreateRefreshHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value(middleware.JWTToken("token"))

		if tkn == nil || tkn == "" {
			logger.Log("Invalid token - ", logger.ERROR)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var req session_dto.RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Log("Invalid request, failed to decode body - "+err.Error(), logger.ERROR)
			http.Error(w, "Invalid request format", http.StatusBadRequest)
			return
		}

		sTkn, ok := tkn.(string)
		if !ok {
			logger.Log("Invalid token", logger.ERROR)
			http.Error(w, "Failed to parse token", int(logger.ERROR))
			return
		}

		usr_id, err := p_db.JWTGen.GetUserIDFromToken(sTkn)
		if err != nil {
			logger.Log("Failed to get userId - "+err.Error(), logger.ERROR)
			http.Error(w, "Failed to get userId", http.StatusUnauthorized)

		}

		valid, err := services.ValidateRefreshToken(p_db, usr_id, req.RefreshToken)
		if err != nil || !valid {
			logger.Log("Invalid refresh token", logger.ERROR)
			http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
			return
		}
		services.RevokeAllSessionTokensToUser(p_db.Redis, usr_id)

		loginData, err := services.GenerateTokensAndSave(p_db, usr_id)
		if err != nil {
			logger.Log("Failed to generate tokens", logger.ERROR)
			http.Error(w, "Failed to generate tokens", http.StatusUnauthorized)
			return
		}

		res := session_dto.RefreshResponse{
			Message:      "Refreshed tokens",
			Token:        loginData.AccessToken,
			RefreshToken: loginData.RefreshToken,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)

	})
}
