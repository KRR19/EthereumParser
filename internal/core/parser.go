package core

import (
	"context"
	"fmt"

	"github.com/KRR19/EthereumParser/internal/models"
	"github.com/KRR19/EthereumParser/pkg/hex"
)

type ParserService struct {
	blockStore BlockStore
}

func NewParserService(blockStore BlockStore) *ParserService {
	return &ParserService{
		blockStore: blockStore,
	}
}

func (p *ParserService) GetCurrentBlock(ctx context.Context) (int, error) {
	n, err := hex.ToDec(p.blockStore.GetLatestBlockNumber())
	if err != nil {
		return 0, fmt.Errorf("failed to convert hex to decimal: %w", err)
	}

	return n, nil
}

func (p *ParserService) Subscribe(ctx context.Context, address string) (bool, error) {
	return true, nil
}

func (p *ParserService) GetTransactions(ctx context.Context, address string) ([]models.Transaction, error) {
	return nil, nil
}
