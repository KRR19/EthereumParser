package main

import (
	"log"
	"net/http"

	"github.com/KRR19/EthereumParser/internal/core"
	"github.com/KRR19/EthereumParser/internal/infrastructure/api"
	"github.com/KRR19/EthereumParser/internal/infrastructure/ethereum"
)

func main() {
	// Initialize components
	parser := core.NewParserService(ethereum.NewClient())
	handler := api.NewHandler(parser)

	// Setup HTTP server
	mux := handler.SetupRoutes()

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
