package core

import (
	"context"
	"testing"
	"time"

	"github.com/KRR19/EthereumParser/internal/core/mock"
)

func TestCrawlerStartStop(t *testing.T) {
	eth := &mock.MockEthereum{}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	crawler.Start(ctx)

	time.Sleep(2 * time.Second)

	cancel()
	crawler.Stop()
}
