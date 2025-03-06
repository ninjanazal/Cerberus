package session_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/dto/session_dto"
	"cerberus/internal/services"
	"cerberus/internal/tools/logger"
	"encoding/json"
	"net/http"
)

// CreateLogoutHandler returns an HTTP handler function for processing logout requests.
// It validates the user's token, ensures the token is active, revokes all session tokens
// for the user, and returns a success response upon successful logout.
//
// Parameters:
//   - p_db: A pointer to the database references (`database.DataRefs`) containing
//     necessary connections and utilities (e.g., JWT validation, Redis, Postgres).
//
// Returns:
//   - http.HandlerFunc: A function that handles the HTTP request and response for logout.
func CreateLogoutHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("token")

		if tkn == nil || tkn == "" {
			logger.Log("Invalid token - ", logger.ERROR)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		sTkn, ok := tkn.(string)
		if !ok {
			logger.Log("Invalid token", logger.ERROR)
			http.Error(w, "Failed to parse token", int(logger.ERROR))
			return
		}

		claims, err := p_db.JWTGen.ValidateJWT(sTkn)
		if err != nil {
			logger.Log("Invalid token - "+err.Error(), logger.ERROR)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		isValid, err := services.IsTokenActive(p_db, claims.UserID, sTkn)
		if err != nil || !isValid {
			logger.Log("Invalid/Revoked token - "+err.Error(), logger.ERROR)
			http.Error(w, "Invalid/Revoked token - "+err.Error(), http.StatusUnauthorized)
			return
		}

		usr, err := services.GetUserById(p_db.Postgres, claims.UserID)
		if err != nil {
			logger.Log("User not found - "+err.Error(), logger.ERROR)
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		services.RevokeAllSessionTokensToUser(p_db.Redis, usr)
		res := session_dto.LogoutResponse{
			Message: "Logged out successfully",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	})
}
