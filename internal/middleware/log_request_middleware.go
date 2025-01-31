package middleware

import (
	logger "cerberus/internal/tools"
	"fmt"
	"net/http"
)

func LogRequestMiddleware(p_next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log(fmt.Sprintf("Request received: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr), logger.INFO)
		p_next.ServeHTTP(w, r)
	})
}
