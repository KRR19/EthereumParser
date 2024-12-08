package api

import (
	"context"

	"github.com/KRR19/EthereumParser/internal/models"
)

type Parser interface {
	GetCurrentBlock(context.Context) (int, error)
	Subscribe(context.Context, string) bool
	GetTransactions(context.Context, string) ([]models.Transaction, error)
}
