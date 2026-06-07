package server

import (
	"net/http"

	"ascii-art-web/internal/handlers"
)

func Start() error {

	http.HandleFunc("/", handlers.HomeHandler)

	return http.ListenAndServe(":8080", nil)
}