package server

import (
	auth_routes "cerberus/internal/routes"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"
	"net/http"
)

func Start() {
	cfg, err := config.LoadEnvFile("../../example.env")

	if err != nil {
		cfg = &config.DefaultCfg
	}

	var mux *http.ServeMux = createRequester()
	auth_routes.SetupAuthRoutes(mux)

	if err := http.ListenAndServe(cfg.GetAddressStr(), mux); err != nil {
		logger.Log.Info().Msg(fmt.Sprintf("Error during serving - %s", err))
	}
}

func createRequester() *http.ServeMux {
	var m *http.ServeMux = http.NewServeMux()

	return m
}
