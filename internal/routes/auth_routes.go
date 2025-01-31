package routes

import (
	"cerberus/internal/handlers"
	"cerberus/internal/middleware"
	logger "cerberus/internal/tools"
	"net/http"
)

func SetupAuthRoutes(p_mux *http.ServeMux) []*Route {
	logger.Log("Setting up Auth Routes", logger.INFO)
	return []*Route{
		NewRoute(p_mux, "/hello", http.HandlerFunc(handlers.HelloWorldHandler), middleware.LogRequestMiddleware),
	}
}
