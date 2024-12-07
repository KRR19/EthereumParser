package core

import (
	"context"
	"time"

	"github.com/KRR19/EthereumParser/internal/models"
)

type Ethereum interface {
	GetLatestBlockNumber(context.Context) (string, error)
	GetTransactionsByBlockNumber(ctx context.Context, blockNumberHex string) ([]models.Transaction, error)
}

type Logger interface {
	Info(string)
	Error(string)
}

type Config interface {
	BlockCheckInterval() time.Duration
	CoreCount() int
}

type BlockStore interface {
	GetLatestBlockNumber() string
	SetBlockNumber(string)
}

type SubscribeStore interface {
	Subscribe(string)
	ValidateTransaction(models.Transaction) bool
	GetSubscribedTransactions(address string) ([]string, bool)
}

type TransactionStore interface {
	Save(models.Transaction)
	GetTransactions(hash ...string) []models.Transaction
}
