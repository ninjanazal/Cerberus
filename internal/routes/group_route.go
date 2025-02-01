package routes

import "net/http"

// GroupRoute represents a group of routes that share a common base path and middlewares.
// All subroutes registered under this group will automatically have the group's base path
// prepended to their own path and will use the group's middlewares in addition to any
// route-specific middlewares.
type GroupRoute struct {
	// mux is the HTTP server multiplexer used to register routes.
	mux *http.ServeMux

	// BaseRoute holds the base path and the shared middlewares for the group.
	BaseRoute
}

// NewGroupRoute creates a new GroupRoute with a specified base path and shared middlewares.
//
// Parameters:
//   - p_mux: A pointer to the http.ServeMux where the routes will be registered.
//   - p_basePath: The base URL path for the route group (e.g., "/auth").
//   - p_middlewares: A variadic slice of Middleware functions that will be applied
//     to all subroutes within this group.
//
// Returns:
//   - A pointer to a newly created GroupRoute.
func NewGroupRoute(p_mux *http.ServeMux, p_basePath string, p_middlewares ...Middleware) *GroupRoute {
	return &GroupRoute{
		mux: p_mux,

		BaseRoute: BaseRoute{
			Path:        p_basePath,
			Middlewares: p_middlewares,
		},
	}
}

// NewRoute creates and registers a new subroute within the GroupRoute.
// The final path of the subroute is the concatenation of the group's base path and the provided subpath.
// The subroute inherits the group's middlewares and may also have additional, route-specific middlewares.
//
// Parameters:
//   - p_subpath: The subpath for the route (e.g., "/login"). It will be appended to the group's base path.
//   - p_handler: The http.Handler that will handle requests for the route.
//   - p_middlewares: A variadic slice of Middleware functions to be applied to this route,
//     in addition to the group's shared middlewares.
//
// Returns:
//   - A pointer to the newly created Route.
func (gRoute *GroupRoute) NewRoute(p_subpath string, p_handler http.Handler, p_middlewares ...Middleware) *Route {
	// Concatenate the group's base path with the subroute's path.
	var p string = gRoute.Path + p_subpath

	// Combine the group's middlewares with the subroute's specific middlewares.
	var mWares []Middleware = append(gRoute.Middlewares, p_middlewares...)

	// Create and register the new route on the group's mux.
	return NewRoute(gRoute.mux, p, p_handler, mWares...)
}
