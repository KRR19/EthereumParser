package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	transactionStore := store.NewTransactionStore()

	crawler := core.NewCrawler(ethereumClient, logger, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	parser := core.NewParserService(blockStore, subscribeStore)
	handler := api.NewHandler(parser)

	mux := handler.SetupRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down server...")

	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Stop the crawler
	crawler.Stop()
	cancel()

	log.Println("Server exiting")
}
