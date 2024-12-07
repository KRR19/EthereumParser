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
}

type BlockStore interface {
	GetLatestBlockNumber() string
	SetBlockNumber(string)
}

type SubscribeStore interface {
	Subscribe(address string)
}
