package ethereum

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestGetLatestBlockNumber_Success(t *testing.T) {
	mockClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       NewMockBody(`{"jsonrpc":"2.0","result":"0x10","id":1}`),
			}, nil
		},
	}

	client := NewClient(EthereumRPCEndpoint, mockClient)
	blockNumber, err := client.GetLatestBlockNumber(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if blockNumber != "0x10" {
		t.Fatalf("Expected block number 0x10, got %s", blockNumber)
	}
}

func TestGetLatestBlockNumber_Failure(t *testing.T) {
	mockClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("network error")
		},
	}

	client := NewClient(EthereumRPCEndpoint, mockClient)
	_, err := client.GetLatestBlockNumber(context.Background())
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestGetTransactionsByBlockNumber_Success(t *testing.T) {
	mockClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       NewMockBody(`{"jsonrpc":"2.0","result":{"transactions":[{"hash":"0x1"}]},"id":1}`),
			}, nil
		},
	}

	client := NewClient(EthereumRPCEndpoint, mockClient)
	transactions, err := client.GetTransactionsByBlockNumber(context.Background(), "0x10")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(transactions) != 1 || transactions[0].Hash != "0x1" {
		t.Fatalf("Expected transaction with hash 0x1, got %v", transactions)
	}
}

func TestGetTransactionsByBlockNumber_Failure(t *testing.T) {
	mockClient := &MockHttpClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("network error")
		},
	}

	client := NewClient(EthereumRPCEndpoint, mockClient)
	_, err := client.GetTransactionsByBlockNumber(context.Background(), "0x10")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
