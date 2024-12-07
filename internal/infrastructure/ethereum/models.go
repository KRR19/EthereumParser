package ethereum

import (
	"github.com/KRR19/EthereumParser/internal/models"
)

type RPCRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Block struct {
	Number       string               `json:"number"`
	Hash         string               `json:"hash"`
	Timestamp    string               `json:"timestamp"`
	Transactions []models.Transaction `json:"transactions"`
}
