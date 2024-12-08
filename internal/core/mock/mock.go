package mock

import (
	"context"
	"errors"
	"time"

	"github.com/KRR19/EthereumParser/internal/models"
)

type MockEthereum struct {
	ShouldFail bool
}

func (m *MockEthereum) GetLatestBlockNumber(ctx context.Context) (string, error) {
	if m.ShouldFail {
		return "", errors.New("failed to get latest block number")
	}
	return "0x10", nil
}

func (m *MockEthereum) GetTransactionsByBlockNumber(ctx context.Context, blockNumberHex string) ([]models.Transaction, error) {
	if m.ShouldFail {
		return nil, errors.New("failed to get transactions by block number")
	}
	return []models.Transaction{{Hash: "0x1"}}, nil
}

type MockLogger struct{}

func (m *MockLogger) Info(msg string)  {}
func (m *MockLogger) Error(msg string) {}

type MockConfig struct{}

func (m *MockConfig) BlockCheckInterval() time.Duration {
	return 1 * time.Second
}

func (m *MockConfig) CoreCount() int {
	return 1
}

type MockBlockStore struct {
	BlockNumber string
}

func (m *MockBlockStore) GetLatestBlockNumber() string {
	return m.BlockNumber
}

func (m *MockBlockStore) SetBlockNumber(blockNumber string) {
	m.BlockNumber = blockNumber
}

type MockSubscribeStore struct {
	Data map[string][]string
}

func (m *MockSubscribeStore) Subscribe(address string) {
	m.Data[address] = []string{}
}

func (m *MockSubscribeStore) ValidateTransaction(tx models.Transaction) bool {
	return tx.To == "0x123"
}

func (m *MockSubscribeStore) GetSubscribedTransactions(address string) ([]string, bool) {
	v, ok := m.Data[address]
	return v, ok
}

type MockTransactionStore struct {
	Transactions map[string]models.Transaction
}

func (m *MockTransactionStore) Save(tx models.Transaction) {
	m.Transactions[tx.Hash] = tx
}

func (m *MockTransactionStore) GetTransactions(hash ...string) []models.Transaction {
	transactionsList := []models.Transaction{}
	for _, h := range hash {
		transactionsList = append(transactionsList, m.Transactions[h])
	}
	return transactionsList
}
