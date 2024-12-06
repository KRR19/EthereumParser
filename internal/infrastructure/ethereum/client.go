package ethereum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ethereumRPCEndpoint = "https://ethereum-rpc.publicnode.com"
)

// Client represents an Ethereum JSON-RPC client
type Client struct {
	endpoint   string
	httpClient *http.Client
}

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
	Number       string   `json:"number"`
	Hash         string   `json:"hash"`
	Timestamp    string   `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

// NewClient creates a new Ethereum JSON-RPC client
func NewClient() *Client {
	return &Client{
		endpoint:   ethereumRPCEndpoint,
		httpClient: &http.Client{},
	}
}

// Call makes a JSON-RPC call to the Ethereum node
func (c *Client) Call(method string, params []interface{}) (*RPCResponse, error) {
	request := RPCRequest{
		JsonRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(c.endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	var response RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if response.Error != nil {
		return nil, fmt.Errorf("RPC error: %s", response.Error.Message)
	}

	return &response, nil
}

// GetLatestBlockNumber retrieves the latest block number
func (c *Client) GetLatestBlockNumber() (uint64, error) {
	response, err := c.Call("eth_blockNumber", nil)
	if err != nil {
		return 0, err
	}

	blockNumberHex, ok := response.Result.(string)
	if !ok {
		return 0, fmt.Errorf("invalid block number format")
	}

	var blockNumber uint64
	fmt.Sscanf(blockNumberHex[2:], "%x", &blockNumber)
	return blockNumber, nil
}

// GetBlockByNumber retrieves a block by its number
func (c *Client) GetBlockByNumber(number uint64) (*Block, error) {
	blockNumberHex := fmt.Sprintf("0x%x", number)
	response, err := c.Call("eth_getBlockByNumber", []interface{}{blockNumberHex, true})
	if err != nil {
		return nil, err
	}

	blockData, err := json.Marshal(response.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal block data: %w", err)
	}

	var block Block
	if err := json.Unmarshal(blockData, &block); err != nil {
		return nil, fmt.Errorf("failed to unmarshal block: %w", err)
	}

	return &block, nil
}
