package auth_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/dto/auth_dto"
	"cerberus/internal/services"
	"cerberus/internal/tools/logger"
	"encoding/json"
	"net/http"
)

// CreateChangePwdHandler creates an HTTP handler for changing a user's password.
//
// This function returns an http.HandlerFunc that processes password change requests.
// The handler performs the following steps:
// 1. Decodes the JSON request body into a ChangePasswordRequest struct
// 2. Calls the ChangePassword service function to process the request
// 3. Responds with a success message or an error
//
// Parameters:
//   - p_db: A pointer to a database.Databases struct, which should contain
//     a Postgres database connection.
//
// Returns:
//   - http.HandlerFunc: A handler function that can be registered with an HTTP server.
//
// Note: This handler uses the logger package for error logging.
func CreateChangePwdHandler(p_db *database.Databases) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req auth_dto.ChangePasswordRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			msg := "Invalid request, failed on decode body - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		err := services.ChangePassword(p_db.Postgres, &req)
		if err != nil {
			logger.Log("Failed to change password - "+err.Error(), logger.ERROR)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res := auth_dto.ChangePasswordResponse{
			Message: "Password changed successfully",
		}
		json.NewEncoder(w).Encode(res)
	})
}
