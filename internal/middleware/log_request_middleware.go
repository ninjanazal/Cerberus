package middleware

import (
	logger "cerberus/internal/tools"
	"fmt"
	"net/http"
)

// LogRequestMiddleware is an HTTP middleware that logs incoming requests.
// It logs the HTTP method, request path, and client IP address before passing
// the request to the next handler.
//
// Parameters:
//   - p_next (http.Handler): The next HTTP handler in the middleware chain.
//
// Returns:
//   - http.Handler: A wrapped HTTP handler that logs request details before processing.
func LogRequestMiddleware(p_next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log(fmt.Sprintf("Request received: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr), logger.INFO)
		p_next.ServeHTTP(w, r)
	})
}
