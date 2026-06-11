package server

import (
	"net/http"

	"ascii-art-web/internal/handlers"
)

func Start() error {
	// map the website routes to their specific handler functions
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	// serve CSS and static assets from the "static" folder
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	// keep the server running and listen on port 8080
	return http.ListenAndServe(":8080", nil)
}
