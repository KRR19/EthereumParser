package store

import "sync"

type BlockStore struct {
	data string
	rwm  sync.RWMutex
}

func NewBlockStore() *BlockStore {
	return &BlockStore{
		data: "",
		rwm:  sync.RWMutex{},
	}
}

func (bs *BlockStore) GetLatestBlockNumber() string {
	bs.rwm.Lock()
	defer bs.rwm.Unlock()
	return bs.data
}

func (bs *BlockStore) SetBlockNumber(data string) {
	bs.rwm.Lock()
	defer bs.rwm.Unlock()
	bs.data = data
}
