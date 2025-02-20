package routes

import (
	"cerberus/internal/handlers"
	"cerberus/internal/middleware"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"net/http"
)

func SetupAuthRoutes(p_mux *http.ServeMux, p_cfg *config.ConfigData) []*Route {
	logger.Log("Setting up Auth Routes", logger.INFO)

	var authGroup GroupRoute = *NewGroupRoute(p_mux, "/auth",
		middleware.TimeRequestMiddleware,
		middleware.CORSMiddleware(p_cfg),
		middleware.LogRequestMiddleware,
	)

	return []*Route{
		NewRoute(p_mux, "/hello",
			http.HandlerFunc(handlers.HelloWorldHandler),
			middleware.TimeRequestMiddleware,
			middleware.CORSMiddleware(p_cfg),
			middleware.LogRequestMiddleware,
		),

		authGroup.NewRoute("/login", http.HandlerFunc(handlers.HelloWorldHandler)),
	}
}
