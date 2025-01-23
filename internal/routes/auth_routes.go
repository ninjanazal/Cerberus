package auth_routes

import (
	logger "cerberus/internal/tools"
	"net/http"
)

func SetupAuthRoutes(*http.ServeMux) {
	logger.Default.Infoln("Setting up Auth Routes")
}
