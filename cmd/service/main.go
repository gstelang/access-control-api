package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gstelang/access-control-api.git/internal/pkg/handlers"
	"github.com/gstelang/access-control-api.git/internal/pkg/middleware"
	"github.com/openfga/go-sdk/client"
	. "github.com/openfga/go-sdk/client"
)

func configureOpenFGA() *client.OpenFgaClient {
	// Configure the OpenFGA client
	fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiUrl: "http://localhost:8080",
		// TODO: change
		StoreId: "01J3YZNXB503K28RM3M7C3625C",
	})
	if err != nil {
		log.Fatalf("Error creating configuration: %v", err)
	}
	return fgaClient
}

func main() {
	fgaClient := configureOpenFGA()
	
	// Set up routes
	http.HandleFunc("/v1/domains/transfer", handlers.TransferHandler)
	// Example http://localhost:8082/v1/authorization/domains/alice.com/check?user=alice&action=can_delete
	http.HandleFunc("/v1/authorization/domains/", middleware.AuthorizationMiddleware(fgaClient)(handlers.DomainAuthorization))

	// Start the server
	fmt.Println("Server starting on port 8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
