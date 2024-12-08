package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/KRR19/EthereumParser/internal/core"
	"github.com/KRR19/EthereumParser/internal/infrastructure/api"
)

func TestServerLifecycle(t *testing.T) {
	deps := initializeDependencies()
	crawler := core.NewCrawler(deps.EthereumClient, deps.Logger, deps.Config, deps.BlockStore, deps.SubscribeStore, deps.TransactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	parser := core.NewParserService(deps.BlockStore, deps.SubscribeStore)
	handler := api.NewHandler(parser)

	server := startServer(handler)

	time.Sleep(time.Second)

	if server.Addr != Addr {
		t.Errorf("Expected server address to be :8080, got %s", server.Addr)
	}

	shutdownServer(server, crawler, cancel)
}

func TestInitializeDependencies(t *testing.T) {
	deps := initializeDependencies()

	if deps.EthereumClient == nil {
		t.Error("Expected EthereumClient to be initialized")
	}
	if deps.Logger == nil {
		t.Error("Expected Logger to be initialized")
	}
	if deps.Config == nil {
		t.Error("Expected Config to be initialized")
	}
	if deps.BlockStore == nil {
		t.Error("Expected BlockStore to be initialized")
	}
	if deps.SubscribeStore == nil {
		t.Error("Expected SubscribeStore to be initialized")
	}
	if deps.TransactionStore == nil {
		t.Error("Expected TransactionStore to be initialized")
	}
}

func TestStartServer(t *testing.T) {
	deps := initializeDependencies()
	crawler := core.NewCrawler(deps.EthereumClient, deps.Logger, deps.Config, deps.BlockStore, deps.SubscribeStore, deps.TransactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	parser := core.NewParserService(deps.BlockStore, deps.SubscribeStore)
	handler := api.NewHandler(parser)

	server := startServer(handler)

	if server.Addr != Addr {
		t.Errorf("Expected server address to be :8080, got %s", server.Addr)
	}

	shutdownServer(server, crawler, cancel)
}

func TestShutdownServer(t *testing.T) {
	deps := initializeDependencies()
	crawler := core.NewCrawler(deps.EthereumClient, deps.Logger, deps.Config, deps.BlockStore, deps.SubscribeStore, deps.TransactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	parser := core.NewParserService(deps.BlockStore, deps.SubscribeStore)
	handler := api.NewHandler(parser)

	server := startServer(handler)

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	shutdownServer(server, crawler, cancel)
}

func TestWaitForShutdown(t *testing.T) {
	deps := initializeDependencies()
	crawler := core.NewCrawler(deps.EthereumClient, deps.Logger, deps.Config, deps.BlockStore, deps.SubscribeStore, deps.TransactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	parser := core.NewParserService(deps.BlockStore, deps.SubscribeStore)
	handler := api.NewHandler(parser)

	server := startServer(handler)

	go func() {
		process, _ := os.FindProcess(os.Getpid())
		process.Signal(os.Interrupt)
	}()

	waitForShutdown(server, crawler, cancel)

	select {
	case <-ctx.Done():
	default:
		t.Error("Expected context to be canceled")
	}
}

func TestRun(t *testing.T) {
	deps := initializeDependencies()
	go func() {
		time.Sleep(time.Second)
		process, _ := os.FindProcess(os.Getpid())
		process.Signal(os.Interrupt)
	}()
	run(deps)
}
