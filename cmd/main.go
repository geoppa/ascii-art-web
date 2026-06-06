package main

import (
	"fmt"
	"log"

	// imports Go's built-in web server and networking tools
	"net/http"
)

func main() {
	// defines the port by creating a variable storing the network port address string
	port := ":8080"

	// serve static files from the current directory. maps the root URL path (/) to serve files from the current folder (.)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// print status message to the terminal
	fmt.Printf("Server starting at http://localhost%s\n", port)

	// starts the network server and listens for incoming browser requests
	err := http.ListenAndServe(port, nil)
	if err != nil {
		// logs the error and immediately shuts down the program
		log.Fatal("Server failed to start: ", err)
	}
}
