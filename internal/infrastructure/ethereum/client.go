package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client represents an Ethereum JSON-RPC client
type Client struct {
	httpClient *http.Client
}

// Call makes a JSON-RPC call to the Ethereum node
func (c *Client) Call(ctx context.Context, method string, params ...interface{}) (*RPCResponse, error) {
	request := RPCRequest{
		JsonRPC: JsonRPC,
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", ethereumRPCEndpoint, bytes.NewReader(body))
	resp, err := c.httpClient.Do(req)
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
func (c *Client) GetLatestBlockNumber(ctx context.Context) (string, error) {
	response, err := c.Call(ctx, "eth_blockNumber")
	if err != nil {
		return "", err
	}

	blockNumberHex, ok := response.Result.(string)
	if !ok {
		return "", fmt.Errorf("invalid block number format")
	}

	return blockNumberHex, nil
}

// GetBlockByNumber retrieves a block by its number
func (c *Client) GetBlockByNumber(ctx context.Context, blockNumberHex string) (*Block, error) {
	response, err := c.Call(ctx, "eth_getBlockByNumber", []interface{}{blockNumberHex, true})
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
