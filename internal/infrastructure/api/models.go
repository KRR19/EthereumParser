package api

import "github.com/KRR19/EthereumParser/internal/models"

type SubscribeReq struct {
	Address string `json:"address"`
}

type GetTransactionsResp struct {
	Transactions []models.Transaction `json:"transactions"`
}
