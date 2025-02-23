package routes

import (
	"cerberus/internal/database"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"
	"net/http"
)

func SetupRoutes(p_mux *http.ServeMux, p_cfg *config.ConfigData, p_dbs *database.Databases) {
	logger.Log("Setting up Routes", logger.INFO)

	var routes []*Route = make([]*Route, 0)
	routes = append(routes, SetupAuthRoutes(p_mux, p_cfg)...)

	listRoutes(routes)
}

func listRoutes(p_routes []*Route) {
	for _, rt := range p_routes {
		logger.Log(fmt.Sprintf("Route added: %s", rt.Path), logger.INFO)
	}
}
