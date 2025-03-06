package routes

import (
	"cerberus/internal/database"
	session_handler "cerberus/internal/handlers/session"
	md "cerberus/internal/middleware"
	"cerberus/internal/tools/logger"
	"cerberus/pkg/config"
	"net/http"
)

// SetupSessionRoutes configures and sets up the session-related routes for the application.
//
// This function creates a group route for session-related endpoints and adds specific routes
// within this group. It applies common middleware to the group and specific middleware to
// individual routes as needed.
//
// Parameters:
//   - p_mux: A pointer to the http.ServeMux to which the routes will be added.
//   - p_cfg: A pointer to the ConfigData structure containing application configuration.
//   - p_dgs: A pointer to the DataRefs structure containing database references.
//
// Returns:
//   - []*Route: A slice of pointers to Route structures representing the configured routes.
func SetupSessionRoutes(p_mux *http.ServeMux, p_cfg *config.ConfigData, p_dgs *database.DataRefs) []*Route {
	logger.Log("📋 Settings up Session Routes", logger.INFO)

	var sessionGroup *GroupRoute = NewGroupRoute(p_mux, "/session",
		md.TimeRequestMiddleware, md.CORSMiddleware(p_cfg), md.LogRequestMiddleware)

	return []*Route{
		sessionGroup.NewRoute("/login", session_handler.CreateLoginHandler(p_dgs),
			md.PostMethodCheckMiddleware),

		sessionGroup.NewRoute("/logout", session_handler.CreateLogoutHandler(p_dgs),
			md.GetMethodCheckMiddleware, md.AuthenticationHeaderMiddleware),
	}
}
