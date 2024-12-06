package core

import (
	"context"

	"github.com/KRR19/EthereumParser/internal/models"
)

// Parser defines the interface for the Ethereum transaction parser
type Parser interface {
	// GetCurrentBlock returns the last parsed block number
	GetCurrentBlock(context.Context) (int, error)

	// Subscribe adds an address to the observer list
	Subscribe(context.Context, string) (bool, error)

	// GetTransactions returns a list of transactions for the given address
	GetTransactions(context.Context, string) ([]models.Transaction, error)
}

// ParserService implements the Parser interface
type ParserService struct {
}

// NewParserService creates a new instance of ParserService
func NewParserService() *ParserService {
	return &ParserService{}
}

// GetCurrentBlock implements Parser.GetCurrentBlock
func (p *ParserService) GetCurrentBlock(ctx context.Context) (int, error) {
	return 1, nil
}

// Subscribe implements Parser.Subscribe
func (p *ParserService) Subscribe(ctx context.Context, address string) (bool, error) {
	return true, nil
}

// GetTransactions implements Parser.GetTransactions
func (p *ParserService) GetTransactions(ctx context.Context, address string) ([]models.Transaction, error) {
	return nil, nil
}
