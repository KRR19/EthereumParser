package store

import (
	"sync"

	"github.com/KRR19/EthereumParser/internal/models"
)

type SubscribeStore struct {
	data map[string][]string
	rwm  *sync.RWMutex
}

func NewSubscribeStore() *SubscribeStore {
	return &SubscribeStore{
		data: make(map[string][]string),
		rwm:  &sync.RWMutex{},
	}
}

func (s *SubscribeStore) Subscribe(address string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.data[address] = []string{}
}

func (s *SubscribeStore) ValidateTransaction(transaction models.Transaction) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	var wasAdded bool
	if _, ok := s.data[transaction.To]; ok {
		s.data[transaction.To] = append(s.data[transaction.To], transaction.Hash)
		wasAdded = true
	}

	if _, ok := s.data[transaction.From]; ok {
		s.data[transaction.From] = append(s.data[transaction.From], transaction.Hash)
		wasAdded = true
	}

	return wasAdded
}

func (s *SubscribeStore) GetSubscribedTransactions(address string) ([]string, bool) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	result, ok := s.data[address]

	return result, ok
}
