package core

import (
	"context"
	"time"
)

type Ethereum interface {
	GetLatestBlockNumber(context.Context) (string, error)
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
