package core

import (
	"context"
	"time"

	"github.com/KRR19/EthereumParser/internal/models"
)

type Crawler struct {
	eth              Ethereum
	log              Logger
	cfg              Config
	blockStore       BlockStore
	subscribeStore   SubscribeStore
	transactionStore TransactionStore
}

func NewCrawler(eth Ethereum, log Logger, cfg Config, blockStore BlockStore, subscribeStore SubscribeStore, transactionStore TransactionStore) *Crawler {
	return &Crawler{
		eth:              eth,
		log:              log,
		cfg:              cfg,
		blockStore:       blockStore,
		subscribeStore:   subscribeStore,
		transactionStore: transactionStore,
	}
}

func (c *Crawler) Start(ctx context.Context) {
	go func() {
		cnt := c.cfg.CoreCount()

		blockNumberChn := make(chan string, cnt)
		go c.fetchBlockNumber(ctx, blockNumberChn)

		newBlockSignal := make(chan string, cnt)
		go c.handleBlockNumber(ctx, blockNumberChn, newBlockSignal)

		transactionChn := make(chan models.Transaction, cnt)
		go c.getBlock(ctx, newBlockSignal, transactionChn)

		for range cnt {
			go c.handleTransaction(ctx, transactionChn)
		}

		<-ctx.Done()
		c.log.Info("Context canceled in Start method")
	}()
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
				newBlockSignal <- bn
				c.blockStore.SetBlockNumber(bn)
			}
		}
	}
}

func (c *Crawler) getBlock(ctx context.Context, newBlockSignal chan string, transactionChn chan models.Transaction) {
	for {
		select {
		case <-ctx.Done():
			c.log.Info("Context canceled in getBlock method")
			return
		case bn, ok := <-newBlockSignal:
			if !ok {
				c.log.Info("New block signal channel closed")
				return
			}

			txs, err := c.eth.GetTransactionsByBlockNumber(ctx, bn)
			if err != nil {
				c.log.Error("Error fetching block: " + err.Error())
				continue
			}

			for _, tx := range txs {
				transactionChn <- tx
			}
		}
	}
}

func (c *Crawler) handleTransaction(ctx context.Context, transactionChn chan models.Transaction) {
	for {
		select {
		case <-ctx.Done():
			c.log.Info("Context canceled in handleTransaction method")
			return
		case tx, ok := <-transactionChn:
			if !ok {
				c.log.Info("Transaction channel closed")
				return
			}

			if c.subscribeStore.ValidateTransaction(tx) {
				c.transactionStore.Save(tx)
			}
		}
	}
}
