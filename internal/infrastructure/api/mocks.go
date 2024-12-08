package api

import (
	"context"
	"errors"

	"github.com/KRR19/EthereumParser/internal/core"
	"github.com/KRR19/EthereumParser/internal/models"
)

type MockParserService struct {
	notFound   bool
	shouldFail bool
}

func (m *MockParserService) GetCurrentBlock(ctx context.Context) (int, error) {
	if m.shouldFail {
		return 0, errors.New("failed to get current block")
	}
	return 10, nil
}

func (m *MockParserService) Subscribe(ctx context.Context, address string) bool {
	return true
}

func (m *MockParserService) GetTransactions(ctx context.Context, address string) ([]models.Transaction, error) {
	if m.shouldFail {
		return nil, errors.New("failed to get transactions")
	}
	if m.notFound {
		return nil, core.ErrAddressNotSubscribed
	}
	return []models.Transaction{}, nil
}
