package core

import (
	"context"
	"time"
)

type Crawler struct {
	eth        Ethereum
	log        Logger
	cfg        Config
	blockStore BlockStore
}

func NewCrawler(eth Ethereum, log Logger, cfg Config, blockStore BlockStore) *Crawler {
	return &Crawler{
		eth:        eth,
		log:        log,
		cfg:        cfg,
		blockStore: blockStore,
	}
}

func (c *Crawler) Start(ctx context.Context) {
    go func() {
        blockNumberChn := make(chan string, 1)
        go c.fetchBlockNumber(ctx, blockNumberChn)

        newBlockSignal := make(chan string, 1)
        go c.handleBlockNumber(ctx, blockNumberChn, newBlockSignal)

        <-ctx.Done()
        c.log.Info("Context canceled in Start method")
    }()
}

func (c *Crawler) handleBlockNumber(ctx context.Context, blockNumberChn chan string, newBlockSignal chan string) {
    defer close(newBlockSignal)
    for {
        select {
        case <-ctx.Done():
            c.log.Info("Context canceled in handleBlockNumber method")
            return
        case bn, ok := <-blockNumberChn:
            if !ok {
                c.log.Info("Block number channel closed")
                return
            }
            localBlock := c.blockStore.GetLatestBlockNumber()
            if bn != localBlock {
                c.log.Info("Fetched block number: " + bn)
                // newBlockSignal <- bn
                c.blockStore.SetBlockNumber(bn)
            }
        }
    }
}

func (c *Crawler) fetchBlockNumber(ctx context.Context, blockNumberChn chan string) {
	defer close(blockNumberChn)
	ticker := time.NewTicker(c.cfg.BlockCheckInterval())
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			c.log.Info("Context canceled in fetchBlockNumbers method")
			return
		case <-ticker.C:
			bn, err := c.eth.GetLatestBlockNumber(ctx)
			if err != nil {
				c.log.Error("Error fetching block number: " + err.Error())
				continue
			}

			blockNumberChn <- bn
		}
	}
}
