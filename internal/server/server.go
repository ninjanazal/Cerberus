package server

import (
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

	var server *http.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: createHandler(),
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Default.Errorln(fmt.Sprintf("Error during serving - %s", err))
	}
}

func createHandler() *http.ServeMux {
	var m *http.ServeMux = http.NewServeMux()

	return m
}
