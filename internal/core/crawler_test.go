package core

import (
	"context"
	"testing"
	"time"

	"github.com/KRR19/EthereumParser/internal/core/mock"
	"github.com/KRR19/EthereumParser/internal/models"
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

func TestGetBlock_Success(t *testing.T) {
	eth := &mock.MockEthereum{}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newBlockSignal := make(chan string, 1)
	transactionChn := make(chan models.Transaction, 1)
	newBlockSignal <- "0x10"

	go crawler.getBlock(ctx, newBlockSignal, transactionChn)

	select {
	case tx := <-transactionChn:
		if tx.Hash != "0x1" {
			t.Fatalf("Expected transaction hash 0x1, got %s", tx.Hash)
		}
	case <-time.After(time.Second):
		t.Fatal("Expected transaction, but got timeout")
	}
}

func TestGetBlock_Failure(t *testing.T) {
	eth := &mock.MockEthereum{ShouldFail: true}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	newBlockSignal := make(chan string, 1)
	transactionChn := make(chan models.Transaction, 1)
	newBlockSignal <- "0x10"

	go crawler.getBlock(ctx, newBlockSignal, transactionChn)

	select {
	case <-transactionChn:
		t.Fatal("Expected no transaction, but got one")
	case <-time.After(time.Second):
	}
}

func TestHandleTransaction_Success(t *testing.T) {
	eth := &mock.MockEthereum{}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{
		Transactions: make(map[string]models.Transaction, 0),
	}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	transactionChn := make(chan models.Transaction, 1)
	transactionChn <- models.Transaction{Hash: "0x1", To: "0x123"}

	go crawler.handleTransaction(ctx, transactionChn)

	time.Sleep(time.Second)
	if len(transactionStore.Transactions) != 1 {
		t.Fatal("Expected transaction to be saved")
	}

}

func TestHandleTransaction_Failure(t *testing.T) {
	eth := &mock.MockEthereum{}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	transactionChn := make(chan models.Transaction, 1)
	transactionChn <- models.Transaction{Hash: "0x1", To: "0x456"}

	go crawler.handleTransaction(ctx, transactionChn)

	time.Sleep(time.Second)
	if len(transactionStore.Transactions) != 0 {
		t.Fatal("Expected no transaction to be saved")
	}

}

func TestHandleBlockNumber_Success(t *testing.T) {
	eth := &mock.MockEthereum{}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	blockNumberChn := make(chan string, 1)
	newBlockSignal := make(chan string, 1)
	blockNumberChn <- "0x10"

	go crawler.handleBlockNumber(ctx, blockNumberChn, newBlockSignal)

	select {
	case bn := <-newBlockSignal:
		if bn != "0x10" {
			t.Fatalf("Expected block number 0x10, got %s", bn)
		}
	case <-time.After(time.Second):
		t.Fatal("Expected block number, but got timeout")
	}
}

func TestHandleBlockNumber_Failure(t *testing.T) {
	eth := &mock.MockEthereum{}
	log := &mock.MockLogger{}
	cfg := &mock.MockConfig{}
	blockStore := &mock.MockBlockStore{}
	subscribeStore := &mock.MockSubscribeStore{}
	transactionStore := &mock.MockTransactionStore{}

	crawler := NewCrawler(eth, log, cfg, blockStore, subscribeStore, transactionStore)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	blockNumberChn := make(chan string, 1)
	newBlockSignal := make(chan string, 1)
	blockNumberChn <- "0x10"
	blockStore.BlockNumber = "0x10"

	go crawler.handleBlockNumber(ctx, blockNumberChn, newBlockSignal)

	select {
	case <-newBlockSignal:
		t.Fatal("Expected no block number, but got one")
	case <-time.After(time.Second):
	}
}
