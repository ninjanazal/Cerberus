package auth_handler

import (
	"cerberus/internal/database"
	"net/http"
)

func CreateHandler(p_db *database.Databases) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
