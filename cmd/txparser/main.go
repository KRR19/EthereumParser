package main

import (
	"context"
	"log"
	"net/http"

	"github.com/KRR19/EthereumParser/internal/core"
	"github.com/KRR19/EthereumParser/internal/infrastructure/api"
	"github.com/KRR19/EthereumParser/internal/infrastructure/config"
	"github.com/KRR19/EthereumParser/internal/infrastructure/ethereum"
	"github.com/KRR19/EthereumParser/internal/infrastructure/logger"
	"github.com/KRR19/EthereumParser/internal/infrastructure/store"
)

func main() {
	ethereumClient := ethereum.NewClient()
	logger := logger.NewLogger()
	cfg := config.NewConfig()
	blockStore := store.NewBlockStore()
	subscribeStore := store.NewSubscribeStore()

	crawler := core.NewCrawler(ethereumClient, logger, cfg, blockStore)
	crawler.Start(context.Background())

	parser := core.NewParserService(blockStore, subscribeStore)
	handler := api.NewHandler(parser)

	mux := handler.SetupRoutes()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
