package routes

import "net/http"

// Middleware defines a function type that takes an http.Handler and returns a new http.Handler.
// This allows for middleware to wrap the handler with additional functionality.
type Middleware func(http.Handler) http.Handler

// Route represents an HTTP route, including its path, handler, and a list of middlewares.
// A Route is created with a specific handler and associated middlewares that will be applied when
// the route is handled by the server.
type Route struct {
	Path        string       // The URL path that this route will handle (e.g., "/home")
	Handler     http.Handler // The handler function to process requests for this route
	Middlewares []Middleware // A list of middlewares to apply to the route before handling requests
}

// NewRoute creates a new route that is registered with the provided ServeMux.
// The route is initialized with the specified path, handler, and a variadic list of middlewares.
// The middlewares will be applied in sequence to the handler before it is registered with the mux.
// The route is then registered with the ServeMux using the provided path and handler.
func NewRoute(p_mux *http.ServeMux, path string, handler http.Handler, middlewares ...Middleware) *Route {
	// Create a new Route with the provided path, handler, and middlewares
	r := &Route{
		Path:        path,
		Handler:     handler,
		Middlewares: middlewares,
	}
	// Apply all middlewares to the route's handler
	r.applyMiddlewares()
	// Register the route with the provided ServeMux
	p_mux.Handle(r.Path, r.Handler)

	// Return the created Route for potential further use
	return r
}

// applyMiddlewares applies the middlewares to the route's handler in sequence.
// This function wraps the handler with each middleware function, creating a new handler
// each time a middleware is applied. The final handler is set on the Route.
func (r *Route) applyMiddlewares() {
	// Start with the original handler
	handler := r.Handler

	// Apply each middleware to the handler in the order they were provided
	for _, middleware := range r.Middlewares {
		handler = middleware(handler)
	}

	// Set the handler to the final result after applying all middlewares
	r.Handler = handler
}
