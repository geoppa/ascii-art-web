package server

import (
	"net/http"

	"ascii-art-web/internal/handlers"
)

func Start() error {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	return http.ListenAndServe(":8080", nil)
}
