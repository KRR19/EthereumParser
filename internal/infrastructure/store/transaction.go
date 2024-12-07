package store

import (
	"sync"

	"github.com/KRR19/EthereumParser/internal/models"
)

type TransactionStore struct {
	data map[string]models.Transaction
	mu   *sync.RWMutex
}

func NewTransactionStore() *TransactionStore {
	return &TransactionStore{
		data: make(map[string]models.Transaction),
		mu:   &sync.RWMutex{},
	}
}

func (s *TransactionStore) Save(transactions models.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[transactions.Hash] = transactions
}

func (s *TransactionStore) GetTransactions(hash ...string) []models.Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()

	transactions := make([]models.Transaction, 0, len(hash))
	for _, h := range hash {
		if t, ok := s.data[h]; ok {
			transactions = append(transactions, t)
		}
	}

	return transactions
}
