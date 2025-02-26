package auth_handler

import (
	"cerberus/internal/database"
	postgres_service "cerberus/internal/database/postgresSQL/service"
	"cerberus/internal/dto/auth_dto"
	logger "cerberus/internal/tools"
	"encoding/json"
	"net/http"
)

// CreateRegisterHandler returns an http.HandlerFunc for handling user registration requests.
//
// Parameters:
//   - p_db: A pointer to the database.Databases struct containing database connections.
//
// Returns:
//   - http.HandlerFunc: A handler function that processes user registration requests.
//
// The handler expects a JSON payload in the request body and returns a JSON response.
// It uses the provided database connection to perform the user registration operation.
func CreateRegisterHandler(p_db *database.Databases) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse Request
		var req auth_dto.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			msg := "Invalid request, failed on decode body - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		usr, err := postgres_service.RegisterUser(p_db.Postgres, &req)
		if err != nil {
			logger.Log(err.Error(), logger.ERROR)
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		// Build response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		res := auth_dto.RegisterResponse{
			Message: "User registered",
			UserId:  usr.ID.String(),
		}
		json.NewEncoder(w).Encode(res)
	})
}
