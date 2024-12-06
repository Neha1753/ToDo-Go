package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-go/routes"
)

func main() {
	r := routes.Router() // Call the Router function from the router package
	fmt.Println("starting the server on port 9000...")

	// Use log.Fatal to log any errors that occur during the HTTP server startup
	log.Fatal(http.ListenAndServe(":9000", r)) // Start the server on port 9000 and use the router
}
