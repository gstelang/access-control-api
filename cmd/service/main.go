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
	// router or mux used interchangeably
	router    *http.ServeMux
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
		router:    http.NewServeMux(),
		fgaClient: fgaClient,
		port:      port,
	}, nil
}

// configureOpenFGA creates and returns an OpenFGA client
func configureOpenFGA() (*client.OpenFgaClient, error) {
	// TODO: hardcoded
	return client.NewSdkClient(&client.ClientConfiguration{
		ApiUrl:  "http://localhost:8080",
		StoreId: "01J4N365K561ZBQ0GXZENCN6AE",
	})
}

// SetupRoutes configures the routes for the server
func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/v1/domains/transfer", handlers.TransferHandler)
	s.router.HandleFunc("/v1/authorization/domains/{domainName}/check", middleware.AuthorizationMiddleware(s.fgaClient)(handlers.DomainAuthorization))
}

// Start begins listening for incoming requests
func (s *Server) Start() error {
	fmt.Printf("Server starting on port %s...\n", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
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
