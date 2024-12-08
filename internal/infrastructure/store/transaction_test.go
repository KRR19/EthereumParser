
package store

import (
	"testing"

	"github.com/KRR19/EthereumParser/internal/models"
)

func TestTransactionStore_Save(t *testing.T) {
	store := NewTransactionStore()
	tx := models.Transaction{Hash: "tx1"}

	store.Save(tx)

	if len(store.data) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(store.data))
	}

	if store.data["tx1"].Hash != "tx1" {
		t.Errorf("Expected transaction hash to be 'tx1', got %s", store.data["tx1"].Hash)
	}
}

func TestTransactionStore_GetTransactions(t *testing.T) {
	store := NewTransactionStore()
	tx1 := models.Transaction{Hash: "tx1"}
	tx2 := models.Transaction{Hash: "tx2"}
	store.Save(tx1)
	store.Save(tx2)

	txs := store.GetTransactions("tx1", "tx2")

	if len(txs) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(txs))
	}

	if txs[0].Hash != "tx1" || txs[1].Hash != "tx2" {
		t.Errorf("Expected transaction hashes to be 'tx1' and 'tx2', got %s and %s", txs[0].Hash, txs[1].Hash)
	}
}

func TestTransactionStore_GetTransactions_NotFound(t *testing.T) {
	store := NewTransactionStore()
	tx := models.Transaction{Hash: "tx1"}
	store.Save(tx)

	txs := store.GetTransactions("tx2")

	if len(txs) != 0 {
		t.Errorf("Expected 0 transactions, got %d", len(txs))
	}
}