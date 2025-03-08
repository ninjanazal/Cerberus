package session_handler

import (
	"cerberus/internal/database"
	"cerberus/internal/middleware"
	"cerberus/internal/tools/logger"
	"net/http"
)

func CreateRefreshHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value(middleware.JWTToken("token"))

		if tkn == nil || tkn == "" {
			logger.Log("Invalid token - ", logger.ERROR)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}
