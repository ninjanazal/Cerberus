package session_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/dto/session_dto"
	"cerberus/internal/services"
	"cerberus/internal/tools/logger"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateLoginHandler creates an HTTP handler function for user login.
//
// This handler performs the following steps:
// 1. Decodes the login request from the request body.
// 2. Authenticates the user using the provided credentials.
// 3. Revokes any existing session tokens for the user.
// 4. Generates new JWT and refresh tokens for the user.
// 5. Responds with the new tokens and a success message.
//
// Parameters:
//   - p_db: A pointer to the DataRefs structure containing database and configuration references.
//
// Returns:
//   - http.HandlerFunc: An HTTP handler function that processes login requests.
//
// The handler responds with different HTTP status codes based on the outcome:
//   - 201 (StatusCreated): Successful login
//   - 409 (StatusConflict): Invalid request body
//   - 401 (StatusUnauthorized): Invalid credentials
//   - 500 (StatusInternalServerError): Server-side error during login process
func CreateLoginHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req session_dto.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			msg := "Invalid request, failed on decode body - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		usr, err := services.AuthenticateUser(p_db.Postgres, &req)
		if err != nil {
			msg := "Invalid credentials - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		services.RevokeAllSessionTokensToUser(p_db.Redis, usr.ID.String())

		loginData, err := services.LoginUser(p_db, usr)
		if err != nil {
			logger.Log("Failed to login user - "+err.Error(), logger.ERROR)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := session_dto.LoginResponse{
			Message:      fmt.Sprintf("%s logged in", usr.Name),
			Token:        loginData.AccessToken,
			RefreshToken: loginData.RefreshToken,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	})
}
