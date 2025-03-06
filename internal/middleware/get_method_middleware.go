package middleware

import (
	"cerberus/internal/tools/logger"
	"net/http"
)

// GetMethodCheckMiddleware is an HTTP middleware that ensures incoming requests
// use the GET method. If the request method is not GET, it responds with a
// "Method not allowed" error and logs the event.
//
// Parameters:
//   - p_next: The next http.Handler in the middleware chain to be called if the request method is GET.
//
// Returns:
//   - http.Handler: A middleware function that processes the request and passes it to the next handler if the method is valid.
func GetMethodCheckMiddleware(p_next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			msg := "Method not allowed"
			logger.Log(msg, logger.ERROR)
			http.Error(w, msg, http.StatusMethodNotAllowed)
			return
		}
		p_next.ServeHTTP(w, r)
	})
}
