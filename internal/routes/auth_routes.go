package auth_routes

import (
	logger "cerberus/internal/tools"
	"net/http"
)

func SetupAuthRoutes(*http.ServeMux) {
	logger.Log.Info().Msg("Setting up Auth Routes")
}
