package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/KRR19/EthereumParser/internal/core"
	"github.com/KRR19/EthereumParser/internal/models"
)

type MockParserService struct {
	notFound   bool
	shouldFail bool
}

func (m *MockParserService) GetCurrentBlock(ctx context.Context) (int, error) {
	if m.shouldFail {
		return 0, errors.New("failed to get current block")
	}
	return 10, nil
}

func (m *MockParserService) Subscribe(ctx context.Context, address string) bool {
	return true
}

func (m *MockParserService) GetTransactions(ctx context.Context, address string) ([]models.Transaction, error) {
	if m.shouldFail {
		return nil, errors.New("failed to get transactions")
	}
	if m.notFound {
		return nil, core.ErrAddressNotSubscribed
	}
	return []models.Transaction{}, nil
}

func TestGetCurrentBlock(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/block", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetCurrentBlock(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"currentBlock":10}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		req, err := http.NewRequest("POST", "/api/v1/block", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetCurrentBlock(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})

	t.Run("internal server error", func(t *testing.T) {
		parser := &MockParserService{shouldFail: true}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/block", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetCurrentBlock(rr, req)

		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
		}
	})
}

func TestSubscribe(t *testing.T) {
	t.Run("successful subscription", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		reqBody := `{"address": "0x123"}`
		req, err := http.NewRequest("POST", "/api/v1/subscribe", strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Subscribe(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"success":true}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/subscribe", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Subscribe(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		reqBody := `{"invalid": "body"}`
		req, err := http.NewRequest("POST", "/api/v1/subscribe", strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Subscribe(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("missing address", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		reqBody := `{"address": ""}`
		req, err := http.NewRequest("POST", "/api/v1/subscribe", strings.NewReader(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.Subscribe(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

func TestGetTransactions(t *testing.T) {
	t.Run("successful request", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/transactions?address=0x123", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetTransactions(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := `{"transactions":[]}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("method not allowed", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		req, err := http.NewRequest("POST", "/api/v1/transactions", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetTransactions(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})

	t.Run("missing address parameter", func(t *testing.T) {
		parser := &MockParserService{}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/transactions", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetTransactions(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}
	})

	t.Run("address not subscribed", func(t *testing.T) {
		parser := &MockParserService{
			notFound: true,
		}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/transactions?address=0x456", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetTransactions(rr, req)

		if status := rr.Code; status != http.StatusForbidden {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusForbidden)
		}
	})

	t.Run("internal server error", func(t *testing.T) {
		parser := &MockParserService{shouldFail: true}
		handler := NewHandler(parser)

		req, err := http.NewRequest("GET", "/api/v1/transactions?address=0x123", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler.GetTransactions(rr, req)

		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
		}
	})
}

func TestSetupRoutes(t *testing.T) {
	parser := &MockParserService{}
	handler := NewHandler(parser)
	mux := handler.SetupRoutes()

	tests := []struct {
		name       string
		method     string
		url        string
		body       string
		wantStatus int
	}{
		{"GetCurrentBlock", "GET", "/api/v1/block", "", http.StatusOK},
		{"GetCurrentBlockMethodNotAllowed", "POST", "/api/v1/block", "", http.StatusMethodNotAllowed},
		{"Subscribe", "POST", "/api/v1/subscribe", `{"address": "0x123"}`, http.StatusOK},
		{"SubscribeMethodNotAllowed", "GET", "/api/v1/subscribe", "", http.StatusMethodNotAllowed},
		{"SubscribeInvalidBody", "POST", "/api/v1/subscribe", `{"invalid": "body"}`, http.StatusBadRequest},
		{"GetTransactions", "GET", "/api/v1/transactions?address=0x123", "", http.StatusOK},
		{"GetTransactionsMethodNotAllowed", "POST", "/api/v1/transactions", "", http.StatusMethodNotAllowed},
		{"GetTransactionsMissingAddress", "GET", "/api/v1/transactions", "", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}
		})
	}
}
