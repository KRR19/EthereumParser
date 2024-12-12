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

const Addr = ":8080"

type Dependencies struct {
	EthereumClient   *ethereum.Client
	Logger           *logger.Logger
	Config           *config.Config
	BlockStore       *store.BlockStore
	SubscribeStore   *store.SubscribeStore
	TransactionStore *store.TransactionStore
}

func main() {
	deps := initializeDependencies()
	run(deps)
}

func run(deps *Dependencies) {
	crawler := core.NewCrawler(deps.EthereumClient, deps.Logger, deps.Config, deps.BlockStore, deps.SubscribeStore, deps.TransactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	parser := core.NewParserService(deps.BlockStore, deps.SubscribeStore, deps.TransactionStore)
	handler := api.NewHandler(parser)

	server := startServer(handler)

	waitForShutdown(server, crawler, cancel)
}

func initializeDependencies() *Dependencies {
	return &Dependencies{
		EthereumClient:   ethereum.NewClient(ethereum.EthereumRPCEndpoint, http.DefaultClient),
		Logger:           logger.NewLogger(os.Stdout),
		Config:           config.NewConfig(),
		BlockStore:       store.NewBlockStore(),
		SubscribeStore:   store.NewSubscribeStore(),
		TransactionStore: store.NewTransactionStore(),
	}
}

func startServer(handler *api.Handler) *http.Server {
	mux := handler.SetupRoutes()
	server := &http.Server{
		Addr:    Addr,
		Handler: mux,
	}
	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	return server
}

func waitForShutdown(server *http.Server, crawler *core.Crawler, cancel context.CancelFunc) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down server...")

	shutdownServer(server, crawler, cancel)

	log.Println("Server exiting")
}

func shutdownServer(server *http.Server, crawler *core.Crawler, cancel context.CancelFunc) {
	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	crawler.Stop()
	cancel()
}
