package auth_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/dto/auth_dto"
	"cerberus/internal/services"
	"cerberus/internal/tools/logger"
	"encoding/json"
	"net/http"
)

// CreateRegisterHandler returns an http.HandlerFunc for handling user registration requests.
//
// Parameters:
//   - p_db: A pointer to the database.DataRefs struct containing database connections and .
//
// Returns:
//   - http.HandlerFunc: A handler function that processes user registration requests.
//
// The handler expects a JSON payload in the request body and returns a JSON response.
// It uses the provided database connection to perform the user registration operation.
func CreateRegisterHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse Request
		var req auth_dto.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			msg := "Invalid request, failed on decode body - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		usr, err := services.RegisterUser(p_db.Postgres, &req)
		if err != nil {
			logger.Log(err.Error(), logger.ERROR)
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		res := auth_dto.RegisterResponse{
			Message: "User registered",
			UserId:  usr.ID.String(),
		}
		json.NewEncoder(w).Encode(res)
	})
}
