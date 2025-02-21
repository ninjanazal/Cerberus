package middleware

import (
	"cerberus/pkg/config"
	"net/http"
	"slices"
)

// CORSMiddleware is a middleware function that handles Cross-Origin Resource Sharing (CORS)
// by inspecting the "Origin" header in incoming HTTP requests. It applies CORS headers
// to the response based on the server's configuration, allowing cross-origin requests
// from approved origins only. This middleware supports the handling of preflight OPTIONS requests.
//
// The CORS policy is controlled via the `p_cfg` configuration object, which contains:
//   - `EnableCORS`: A flag to enable or disable CORS support globally for the server.
//   - `AllowedOrigins`: A map of allowed origins (as keys) for cross-origin requests. Requests from
//     origins not listed in this map will be denied with a "403 Forbidden" error.
//   - If `EnableCORS` is set to `false`, the middleware will not apply any CORS headers and
//     the request will be passed through as normal.
//
// This middleware also supports preflight requests (OPTIONS requests), which are sent by browsers
// to check if the actual request is allowed by the server. If the `OPTIONS` method is detected,
// the server will respond with a `204 No Content` status to indicate that the request is safe to send.
//
// Parameters:
//   - p_cfg: A pointer to a `ConfigData` object, which holds the CORS configuration for the server.
//     The configuration object must contain:
//   - `EnableCORS`: A boolean flag indicating if CORS should be enabled.
//   - `AllowedOrigins`: A map of allowed origins, with the origin as the key and a boolean `true`
//     value indicating that the origin is allowed.
//
// Return:
//   - A function that takes an `http.Handler` and returns a new handler with CORS functionality applied.
//     The returned handler will set the appropriate CORS headers on the response if the origin is allowed
//     and handle preflight OPTIONS requests.
//
// Example usage:
//
//	// Create the CORSMiddleware using the configuration
//	corsMiddleware := middleware.CORSMiddleware(cfg)
//
//	// Apply CORS middleware in the route setup
//	NewRoute(mux, "/example",
//	    http.HandlerFunc(exampleHandler),
//	    corsMiddleware,  // Apply CORS middleware
//	    otherMiddlewares...,
//	)
func CORSMiddleware(p_cfg *config.ConfigData) func(http.Handler) http.Handler {
	return func(p_next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			// If the request has no Origin, it's not from a browser (e.g., mobile app, backend service)
			if !p_cfg.EnableCORS || origin == "" {
				p_next.ServeHTTP(w, r)
				return
			}

			// Check if the origin is allowed
			if !slices.Contains(p_cfg.AllowedOrigins, origin) {
				http.Error(w, "CORS policy does not allow this origin", http.StatusForbidden)
				return
			}

			// Set CORS headers for allowed origins
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// Handle preflight requests (OPTIONS)
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// Call the next handler
			p_next.ServeHTTP(w, r)

		})
	}
}
