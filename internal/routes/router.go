package routes

import (
	logger "cerberus/internal/tools"
	"fmt"
	"net/http"
)

func SetupRoutes(p_mux *http.ServeMux) {
	logger.Log("Setting up Routes", logger.INFO)

	var routes []*Route = make([]*Route, 0)
	routes = append(routes, SetupAuthRoutes(p_mux)...)

	listRoutes(routes)
}

func listRoutes(p_routes []*Route) {
	for _, rt := range p_routes {
		logger.Log(fmt.Sprintf("Route added: %s", rt.Path), logger.INFO)
	}
}
