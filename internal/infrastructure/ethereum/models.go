package ethereum

import (
	"net/http"

	"github.com/KRR19/EthereumParser/internal/models"
)

// RPCRequest represents a JSON-RPC request
type RPCRequest struct {
	JsonRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// RPCResponse represents a JSON-RPC response
type RPCResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

// RPCError represents a JSON-RPC error
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Block represents an Ethereum block
type Block struct {
	Number       string               `json:"number"`
	Hash         string               `json:"hash"`
	Timestamp    string               `json:"timestamp"`
	Transactions []models.Transaction `json:"transactions"`
}

// NewClient creates a new Ethereum JSON-RPC client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}
