package core

import (
	"context"
	"fmt"

	"github.com/KRR19/EthereumParser/internal/models"
	"github.com/KRR19/EthereumParser/pkg/hex"
)

type ParserService struct {
	blockStore     BlockStore
	subscribeStore SubscribeStore
	transactionStore TransactionStore
}

func NewParserService(blockStore BlockStore, subscribeStore SubscribeStore) *ParserService {
	return &ParserService{
		blockStore:     blockStore,
		subscribeStore: subscribeStore,
	}
}

func (p *ParserService) GetCurrentBlock(ctx context.Context) (int, error) {
	n, err := hex.ToDec(p.blockStore.GetLatestBlockNumber())
	if err != nil {
		return 0, fmt.Errorf("failed to convert hex to decimal: %w", err)
	}

	return n, nil
}

func (p *ParserService) Subscribe(ctx context.Context, address string) bool {
	p.subscribeStore.Subscribe(address)
	return true
}

func (p *ParserService) GetTransactions(ctx context.Context, address string) ([]models.Transaction, error) {
	txs, ok := p.subscribeStore.GetSubscribedTransactions(address)
	if ok {
		return nil, ErrAddressNotSubscribed
	}

	return p.transactionStore.GetTransactions(txs...), nil
}
