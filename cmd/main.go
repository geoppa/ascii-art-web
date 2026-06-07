package main

import (
	"fmt"
	"log"

	"ascii-art-web/internal/server"
)

func main() {

	fmt.Println("Server starting at http://localhost:8080")

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}