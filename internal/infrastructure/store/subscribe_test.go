package store

import (
	"testing"

	"github.com/KRR19/EthereumParser/internal/models"
)

func TestSubscribeStore_Subscribe(t *testing.T) {
	store := NewSubscribeStore()
	address := "0x123"

	store.Subscribe(address)

	if _, ok := store.data[address]; !ok {
		t.Errorf("Expected address %s to be subscribed", address)
	}
}

func TestSubscribeStore_ValidateTransaction(t *testing.T) {
	store := NewSubscribeStore()
	address := "0x123"
	store.Subscribe(address)

	tx := models.Transaction{To: address, Hash: "tx1"}
	if !store.ValidateTransaction(tx) {
		t.Errorf("Expected transaction to be validated and added")
	}

	if len(store.data[address]) != 1 || store.data[address][0] != "tx1" {
		t.Errorf("Expected transaction hash to be added to the address")
	}
}

func TestSubscribeStore_GetSubscribedTransactions(t *testing.T) {
	store := NewSubscribeStore()
	address := "0x123"
	store.Subscribe(address)
	store.data[address] = []string{"tx1", "tx2"}

	txs, ok := store.GetSubscribedTransactions(address)
	if !ok {
		t.Errorf("Expected to find subscribed transactions for address %s", address)
	}

	if len(txs) != 2 || txs[0] != "tx1" || txs[1] != "tx2" {
		t.Errorf("Expected to get correct subscribed transactions")
	}
}

func TestSubscribeStore_GetSubscribedTransactions_NotSubscribed(t *testing.T) {
	store := NewSubscribeStore()
	address := "0x123"

	_, ok := store.GetSubscribedTransactions(address)
	if ok {
		t.Errorf("Expected not to find subscribed transactions for address %s", address)
	}
}
