package routes

import "net/http"

// Middleware defines a function type that takes an http.Handler and returns a new http.Handler.
// This allows for middleware to wrap the handler with additional functionality.
type Middleware func(http.Handler) http.Handler

type BaseRoute struct {
	Path        string       // The URL path that this route will handle (e.g., "/home")
	Middlewares []Middleware // A list of middlewares to apply to the route before handling requests
}
