package server

import (
	"cerberus/internal/routes"
	logger "cerberus/internal/tools"
	"cerberus/pkg/config"
	"fmt"
	"net/http"
	"os"
)

func Start() {
	var file_path string = os.Getenv("CONFIG_FILE")
	cfg, err := config.LoadEnvFile(file_path)

	if err != nil {
		cfg = &config.DefaultCfg
	}

	var mux *http.ServeMux = http.NewServeMux()
	routes.SetupRoutes(mux, cfg)

	logger.Log(fmt.Sprintf("🍭 Starting server at %s", cfg.GetAddressStr()), logger.INFO)

	if err := http.ListenAndServe(cfg.GetAddressStr(), mux); err != nil {
		logger.Log(fmt.Sprintf("💥 Error during serving - %s", err), logger.INFO)
	}
}
