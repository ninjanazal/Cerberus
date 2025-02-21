package middleware

import (
	logger "cerberus/internal/tools"
	"fmt"
	"net/http"
	"time"
)

// TimeRequestMiddleware is an HTTP middleware that measures the duration of each request.
// It logs the processing time in DEBUG level using the logger.
//
// Parameters:
//   - next (http.Handler): The next HTTP handler in the middleware chain.
//
// Returns:
//   - http.Handler: A wrapped HTTP handler that measures and logs request duration.
func TimeRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var startT time.Time = time.Now()

		next.ServeHTTP(w, r)
		var d time.Duration = time.Since(startT)
		logger.Log(fmt.Sprintf("Processed in: %s", d), logger.DEBUG)
	})
}
