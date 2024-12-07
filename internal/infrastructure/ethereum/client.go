package ethereum

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KRR19/EthereumParser/internal/models"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ethereumRPCEndpoint, bytes.NewReader(body))
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

func (c *Client) GetTransactionsByBlockNumber(ctx context.Context, blockNumberHex string) ([]models.Transaction, error) {
	response, err := c.Call(ctx, "eth_getBlockByNumber", blockNumberHex, true)
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

	return block.Transactions, nil
}
