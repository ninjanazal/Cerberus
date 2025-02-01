package routes

import (
	"cerberus/internal/handlers"
	"cerberus/internal/middleware"
	logger "cerberus/internal/tools"
	"net/http"
)

func SetupAuthRoutes(p_mux *http.ServeMux) []*Route {
	logger.Log("Setting up Auth Routes", logger.INFO)
	var authGroup GroupRoute = *NewGroupRoute(p_mux, "/auth", middleware.LogRequestMiddleware)

	return []*Route{
		NewRoute(p_mux, "/hello", http.HandlerFunc(handlers.HelloWorldHandler), middleware.LogRequestMiddleware),

		authGroup.NewRoute("/login", http.HandlerFunc(handlers.HelloWorldHandler)),
	}
}
