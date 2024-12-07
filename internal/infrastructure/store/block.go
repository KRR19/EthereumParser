package store

import "sync"

type BlockStore struct {
	data string
	mu   sync.Mutex
}

func NewBlockStore() *BlockStore {
	return &BlockStore{}
}

func (bs *BlockStore) GetLatestBlockNumber() string {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.data
}

func (bs *BlockStore) SetBlockNumber(data string) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.data = data
}
