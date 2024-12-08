package mock

import (
	"context"
	"time"

	"github.com/KRR19/EthereumParser/internal/models"
)

type MockEthereum struct{}

func (m *MockEthereum) GetLatestBlockNumber(ctx context.Context) (string, error) {
	return "0x10", nil
}

func (m *MockEthereum) GetTransactionsByBlockNumber(ctx context.Context, blockNumberHex string) ([]models.Transaction, error) {
	return []models.Transaction{}, nil
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
	return true
}
func (m *MockSubscribeStore) GetSubscribedTransactions(address string) ([]string, bool) {
	v, ok := m.Data[address]
	return v, ok
}

type MockTransactionStore struct {
	Transactions map[string]models.Transaction
}

func (m *MockTransactionStore) Save(tx models.Transaction) {}
func (m *MockTransactionStore) GetTransactions(hash ...string) []models.Transaction {
	transactionsList := []models.Transaction{}
	for _, h := range hash {
		transactionsList = append(transactionsList, m.Transactions[h])
	}

	return transactionsList
}