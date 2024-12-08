package store

import (
	"strconv"
	"sync"
	"testing"
)

func TestBlockStore_GetLatestBlockNumber(t *testing.T) {
	bs := NewBlockStore()
	bs.SetBlockNumber("0x10")

	blockNumber := bs.GetLatestBlockNumber()
	if blockNumber != "0x10" {
		t.Errorf("Expected block number to be 0x10, got %s", blockNumber)
	}
}

func TestBlockStore_SetBlockNumber(t *testing.T) {
	bs := NewBlockStore()
	bs.SetBlockNumber("0x20")

	blockNumber := bs.GetLatestBlockNumber()
	if blockNumber != "0x20" {
		t.Errorf("Expected block number to be 0x20, got %s", blockNumber)
	}
}

func TestBlockStore_Concurrency(t *testing.T) {
	bs := NewBlockStore()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			bs.SetBlockNumber("0x" + strconv.Itoa(i))
		}(i)
	}

	wg.Wait()

	blockNumber := bs.GetLatestBlockNumber()
	if blockNumber == "" {
		t.Errorf("Expected block number to be set, got empty string")
	}
}
