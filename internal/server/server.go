package server

import (
	"net/http"

	"ascii-art-web/internal/handlers"
)

func Start() error {

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	return http.ListenAndServe(":8080", nil)
}