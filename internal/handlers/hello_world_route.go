package handlers

import (
	logger "cerberus/internal/tools"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	logger.Log("On HelloWorldHandler", logger.INFO)

	w.Write([]byte("Hello World"))
}
