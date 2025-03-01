package routes

import (
	"cerberus/internal/database"
	auth_handler "cerberus/internal/handlers/auth"
	md "cerberus/internal/middleware"
	logger "cerberus/internal/tools/logger"
	"cerberus/pkg/config"
	"net/http"
)

// SetupAuthRoutes configures and returns the authentication-related routes for the application.
//
// Parameters:
//   - p_mux: A pointer to the http.ServeMux that will handle the routes.
//   - p_cfg: A pointer to the config.ConfigData containing application configuration.
//   - p_dbs: A pointer to the database.Databases struct for database connections.
//
// Returns:
//   - []*Route: A slice of pointers to Route structs representing the configured routes.
//
// Each route is configured with appropriate handlers and middleware. The function
// uses a GroupRoute for the "/auth" prefix to apply common middleware to all auth routes.
func SetupAuthRoutes(p_mux *http.ServeMux, p_cfg *config.ConfigData, p_dbs *database.Databases) []*Route {
	logger.Log("🔒 Setting up Auth Routes", logger.INFO)

	var authGroup *GroupRoute = NewGroupRoute(p_mux, "/auth",
		md.TimeRequestMiddleware, md.CORSMiddleware(p_cfg), md.LogRequestMiddleware)

	return []*Route{
		authGroup.NewRoute("/register", auth_handler.CreateRegisterHandler(p_dbs),
			md.PostMethodCheckMiddleware),

		authGroup.NewRoute("/change-password", auth_handler.CreateChangePwdHandler(p_dbs),
			md.PostMethodCheckMiddleware),
	}
}
