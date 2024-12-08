package core

import (
	"context"
	"testing"

	"github.com/KRR19/EthereumParser/internal/core/mock"
	"github.com/KRR19/EthereumParser/internal/models"
)

func TestGetCurrentBlock(t *testing.T) {
	blockStore := &mock.MockBlockStore{BlockNumber: "0x10"}
	subscribeStore := &mock.MockSubscribeStore{}
	parser := NewParserService(blockStore, subscribeStore)

	block, err := parser.GetCurrentBlock(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if block != 16 {
		t.Fatalf("Expected block number 16, got %d", block)
	}
}

func TestSubscribe(t *testing.T) {
	subscribeStore := &mock.MockSubscribeStore{Data: map[string][]string{}}
	blockStore := &mock.MockBlockStore{}
	parser := NewParserService(blockStore, subscribeStore)

	address := "0x123"
	subscribed := parser.Subscribe(context.Background(), address)

	if !subscribed {
		t.Fatalf("Expected subscription to be successful")
	}

	if _, ok := subscribeStore.Data[address]; !ok {
		t.Fatalf("Expected address %s to be subscribed", address)
	}
}

func TestGetTransactions(t *testing.T) {
	transactionStore := &mock.MockTransactionStore{
		Transactions: map[string]models.Transaction{
			"tx1": {Hash: "tx1"},
			"tx2": {Hash: "tx2"},
		},
	}
	subscribeStore := &mock.MockSubscribeStore{
		Data: map[string][]string{
			"0x123": {"tx1", "tx2"},
		},
	}
	blockStore := &mock.MockBlockStore{}
	parser := NewParserService(blockStore, subscribeStore)
	parser.transactionStore = transactionStore

	t.Run("address subscribed", func(t *testing.T) {
		txs, err := parser.GetTransactions(context.Background(), "0x123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(txs) != 2 {
			t.Fatalf("Expected 2 transactions, got %d", len(txs))
		}
	})

	t.Run("address not subscribed", func(t *testing.T) {
		_, err := parser.GetTransactions(context.Background(), "0x456")
		if err != ErrAddressNotSubscribed {
			t.Fatalf("Expected ErrAddressNotSubscribed, got %v", err)
		}
	})
}
