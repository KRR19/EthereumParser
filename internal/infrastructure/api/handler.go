package api

import (
	"encoding/json"
	"net/http"

	"github.com/KRR19/EthereumParser/internal/core"
)

type Handler struct {
	parser Parser
}

func NewHandler(parser Parser) *Handler {
	return &Handler{
		parser: parser,
	}
}

func (h *Handler) GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()
	block, err := h.parser.GetCurrentBlock(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(map[string]int{"currentBlock": block}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SubscribeReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	success := h.parser.Subscribe(ctx, req.Address)

	if err := json.NewEncoder(w).Encode(map[string]bool{"success": success}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Address parameter is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	transactions, err := h.parser.GetTransactions(ctx, address)
	if err == core.ErrAddressNotSubscribed {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	resp := GetTransactionsResp{Transactions: transactions}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/block", h.GetCurrentBlock)
	mux.HandleFunc("/api/v1/subscribe", h.Subscribe)
	mux.HandleFunc("/api/v1/transactions", h.GetTransactions)

	return mux
}
