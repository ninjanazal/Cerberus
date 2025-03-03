package session_handler

import (
	"cerberus/internal/database"
	"net/http"
)

func CreateLogoutHandler(p_db *database.DataRefs) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}
