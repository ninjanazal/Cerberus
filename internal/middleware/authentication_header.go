package middleware

import (
	"context"
	"net/http"
	"strings"
)

func authenticationHeaderMiddleware(p_next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "token", strings.TrimPrefix(authHeader, "Bearer "))

		p_next.ServeHTTP(w, r.WithContext(ctx))
	})
}
