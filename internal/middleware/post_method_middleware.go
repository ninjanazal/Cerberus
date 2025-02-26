package middleware

import (
	logger "cerberus/internal/tools"
	"net/http"
)

// PostMethodCheckMiddleware is an HTTP middleware that ensures the incoming request uses the POST method.
//
// Parameters:
//   - p_next: The next http.Handler in the middleware chain.
//
// Returns:
//   - An http.Handler that wraps the provided handler with POST method checking.
func PostMethodCheckMiddleware(p_next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			msg := "Method not allowed"
			logger.Log(msg, logger.INFO)
			http.Error(w, msg, http.StatusMethodNotAllowed)
			return
		}

		p_next.ServeHTTP(w, r)
	})
}
