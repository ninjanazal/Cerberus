package routes

import (
	"cerberus/internal/database"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"
	"net/http"
)

// SetupRoutes configures all the routes for the application.
//
// Parameters:
//   - p_mux: A pointer to the http.ServeMux that will handle the routes.
//   - p_cfg: A pointer to the config.ConfigData containing application configuration.
//   - p_dbs: A pointer to the database.Databases struct for database connections.
//
// The function doesn't return anything, but it modifies the provided ServeMux
// by adding routes to it.
func SetupRoutes(p_mux *http.ServeMux, p_cfg *config.ConfigData, p_dbs *database.Databases) {
	logger.Log("Setting up Routes", logger.INFO)

	var routes []*Route = make([]*Route, 0)
	routes = append(routes, SetupAuthRoutes(p_mux, p_cfg, p_dbs)...)
	routes = append(routes, SetupSessionRoutes(p_mux, p_cfg, p_dbs)...)

	listRoutes(routes)
}

// region Private

// listRoutes logs all the configured routes.
//
// Parameters:
//   - p_routes: A slice of pointers to Route structs representing the configured routes.
//
// This function iterates through the provided routes and logs each route's path.
// It's used for debugging and informational purposes to show which routes have been
// successfully set up in the application.
func listRoutes(p_routes []*Route) {
	for _, rt := range p_routes {
		logger.Log(fmt.Sprintf("Route added: %s", rt.Path), logger.INFO)
	}
}

// endregion Private
