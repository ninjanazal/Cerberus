package handlers

import (
	logger "cerberus/internal/tools"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	logger.Default.Infoln("On HelloWorldHandler")
}
