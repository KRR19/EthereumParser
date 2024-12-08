package main

import (
	"context"
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

	time.Sleep(1 * time.Second)

	if server.Addr != Addr {
		t.Errorf("Expected server address to be :8080, got %s", server.Addr)
	}

	shutdownServer(server, crawler, cancel)
}
