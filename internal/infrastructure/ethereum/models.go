package ethereum

import "net/http"

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
	Number       string        `json:"number"`
	Hash         string        `json:"hash"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
}

// NewClient creates a new Ethereum JSON-RPC client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

type Transaction struct {
	Type                 string   `json:"type"`
	ChainID              string   `json:"chainId"`
	Nonce                string   `json:"nonce"`
	Gas                  string   `json:"gas"`
	MaxFeePerGas         string   `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	To                   string   `json:"to"`
	Value                string   `json:"value"`
	AccessList           []string `json:"accessList"`
	Input                string   `json:"input"`
	R                    string   `json:"r"`
	S                    string   `json:"s"`
	YParity              string   `json:"yParity"`
	V                    string   `json:"v"`
	Hash                 string   `json:"hash"`
	BlockHash            string   `json:"blockHash"`
	BlockNumber          string   `json:"blockNumber"`
	TransactionIndex     string   `json:"transactionIndex"`
	From                 string   `json:"from"`
	GasPrice             string   `json:"gasPrice"`
}
