package server

import (
	"cerberus/internal/database"
	"cerberus/internal/routes"
	logger "cerberus/internal/tools/logger"
	"cerberus/pkg/config"
	"fmt"
	"net/http"
	"os"
)

// Start initializes and runs the server application.

// The function uses environment variables and configuration files to set up
// the application. If there's an error loading the configuration, it falls
// back to default configuration values.
//
// The server address is determined by the configuration, and all major steps
// are logged for monitoring and debugging purposes.
//
// If any critical error occurs during database initialization or server startup,
// the function will log the error and return, effectively stopping the application.
func Start() {
	logger.Log("🃏 Running Cerberus", logger.INFO)

	// Load configuration from environment file
	var file_path string = os.Getenv("CONFIG_FILE")
	cfg, err := config.LoadEnvFile(file_path)

	if err != nil {
		cfg = &config.DefaultCfg
	}

	// Initialize HTTP multiplexer
	var mux *http.ServeMux = http.NewServeMux()

	// Initialize database connections
	dbs, err := database.InitDatabases(cfg)
	if err != nil {
		logger.Log(fmt.Sprintf("Something went wrong! %s", err.Error()), logger.ERROR)
		return
	}

	// Define routes
	routes.SetupRoutes(mux, cfg, dbs)

	logger.Log(fmt.Sprintf("🍭 Starting server at %s", cfg.GetAddressStr()), logger.INFO)
	// Start the HTTP server
	if err := http.ListenAndServe(cfg.GetAddressStr(), mux); err != nil {
		logger.Log(fmt.Sprintf("💥 Error during serving - %s", err), logger.INFO)
	}
}
