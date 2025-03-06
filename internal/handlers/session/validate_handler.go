package session_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/middleware"
	"cerberus/internal/services"
	"cerberus/internal/tools/logger"
	"net/http"
)

// CreateValidateHandler returns an HTTP handler function that validates a JWT token from the request context.
// It checks if the token is present, parses it as a string, and validates it using the provided JWT generator.
// If the token is valid and active, the handler responds with a 200 OK status. Otherwise, it returns an appropriate
// error response with details about the failure.
//
// Parameters:
//   - p_db: A pointer to the database.DataRefs struct containing dependencies such as the JWT generator and token validation service.
//
// Returns:
//   - http.HandlerFunc: A function that handles HTTP requests and validates JWT tokens.
func CreateValidateHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value(middleware.JWTToken("token"))

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
			logger.Log("Revoked token", logger.ERROR)
			http.Error(w, "Revoked token", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
