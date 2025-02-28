package routes

import (
	"cerberus/internal/database"
	session_handler "cerberus/internal/handlers/session"
	md "cerberus/internal/middleware"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"net/http"
)

func SetupSessionRoutes(p_mux *http.ServeMux, p_cfg *config.ConfigData, p_dgs *database.Databases) []*Route {
	logger.Log("📋 Settings up Session Routes", logger.INFO)

	var sessionGroup *GroupRoute = NewGroupRoute(p_mux, "/session",
		md.TimeRequestMiddleware, md.CORSMiddleware(p_cfg), md.LogRequestMiddleware)

	return []*Route{
		sessionGroup.NewRoute("/login", session_handler.CreateLoginHandler(p_dgs),
			md.PostMethodCheckMiddleware),
	}
}
