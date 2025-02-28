package session_handler

import (
	"cerberus/internal/database"
	postgres_service "cerberus/internal/database/postgresSQL/service"
	"cerberus/internal/dto/session_dto"
	logger "cerberus/internal/tools"
	"encoding/json"
	"net/http"
)

func CreateLoginHandler(p_db *database.Databases) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req session_dto.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			msg := "Invalid request, failed on decode body - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		usr, err := postgres_service.LoginUserPasswordCheck(p_db.Postgres, &req)
		if err != nil {
			msg := "Invalid credentials - " + err.Error()
			logger.Log(msg, logger.ERROR)

			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

	})
}
