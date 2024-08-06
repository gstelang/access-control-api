package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gstelang/access-control-api.git/internal/pkg/handlers"
	"github.com/gstelang/access-control-api.git/internal/pkg/middleware"
	"github.com/openfga/go-sdk/client"
)

// Server struct to hold dependencies and configuration
type Server struct {
	fgaClient *client.OpenFgaClient
	port      string
}

// NewServer creates and returns a new Server instance
func NewServer(port string) (*Server, error) {
	fgaClient, err := configureOpenFGA()
	if err != nil {
		return nil, fmt.Errorf("failed to configure OpenFGA: %w", err)
	}

	return &Server{
		fgaClient: fgaClient,
		port:      port,
	}, nil
}

// configureOpenFGA creates and returns an OpenFGA client
func configureOpenFGA() (*client.OpenFgaClient, error) {
	// TODO: hardcoded
	return client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:  "http://localhost:8080",
		StoreId: "01J3YZNXB503K28RM3M7C3625C",
	})
}

// SetupRoutes configures the routes for the server
func (s *Server) SetupRoutes() {
	http.HandleFunc("/v1/domains/transfer", handlers.TransferHandler)
	http.HandleFunc("/v1/authorization/domains/", middleware.AuthorizationMiddleware(s.fgaClient)(handlers.DomainAuthorization))
}

// Start begins listening for incoming requests
func (s *Server) Start() error {
	fmt.Printf("Server starting on port %s...\n", s.port)
	return http.ListenAndServe(":"+s.port, nil)
}

func main() {
	server, err := NewServer("8082")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	server.SetupRoutes()

	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
