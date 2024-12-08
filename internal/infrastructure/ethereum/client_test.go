package ethereum

import (
	"context"
	"net/http"
	"testing"
)

func TestGetLatestBlockNumber(t *testing.T) {
	client := NewClient(EthereumRPCEndpoint, http.DefaultClient)
	blockNumber, err := client.GetLatestBlockNumber(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if blockNumber == "" {
		t.Fatalf("Expected block number, got empty string")
	}
}
