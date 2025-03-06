package middleware

import (
	"context"
	"net/http"
	"strings"
)

// JWTToken is a custom type representing a JSON Web Token (JWT) string.
// It is used to provide type safety and clarity when working with JWTs in the application,
// avoiding potential collisions or misuse of generic string types.
type JWTToken string

// AuthenticationHeaderMiddleware is an HTTP middleware that validates the "Authorization" header
// in incoming requests. It ensures the header is present, checks for the "Bearer " prefix, and
// extracts the token for use in subsequent handlers.
//
// Parameters:
//   - p_next: The next http.Handler in the middleware chain to be called after validation.
//
// Returns:
//   - http.Handler: A middleware function that processes the request and passes it to the next handler.
func AuthenticationHeaderMiddleware(p_next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, JWTToken("token"), strings.TrimPrefix(authHeader, "Bearer "))

		p_next.ServeHTTP(w, r.WithContext(ctx))
	})
}
