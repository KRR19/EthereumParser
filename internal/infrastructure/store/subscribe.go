package store

import "sync"

type SubscribeStore struct {
	data map[string]bool
	rwm  *sync.RWMutex
}

func NewSubscribeStore() *SubscribeStore {
	return &SubscribeStore{
		data: make(map[string]bool),
		rwm:  &sync.RWMutex{},
	}
}

func (s *SubscribeStore) Subscribe(address string) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.data[address] = true
}

func (s *SubscribeStore) IsAddressSubscribed(address string) bool {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.data[address]
}
