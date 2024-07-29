package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gstelang/access-control-api.git/internal/pkg/handlers"
)

func main() {
	// Set up routes
	http.HandleFunc("/v1/domains/transfer", handlers.TransferHandler)
	http.HandleFunc("/v1/domains/editNameServer", handlers.EditNameServerHandler)

	// Start the server
	fmt.Println("Server starting on port 8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
