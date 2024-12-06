package api

import (
	"encoding/json"
	"net/http"

	"github.com/KRR19/EthereumParser/internal/core"
)

// Handler handles HTTP requests
type Handler struct {
	parser core.Parser
}

// NewHandler creates a new HTTP handler
func NewHandler(parser core.Parser) *Handler {
	return &Handler{
		parser: parser,
	}
}

// GetCurrentBlock handles the request to get the current block
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
	}
	json.NewEncoder(w).Encode(map[string]int{"currentBlock": block})
}

// Subscribe handles the request to subscribe to an address
func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Address string `json:"address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Address == "" {
		http.Error(w, "Address is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	success := h.parser.Subscribe(req.Address)
	json.NewEncoder(w).Encode(map[string]bool{"success": success})
}

// GetTransactions handles the request to get transactions for an address
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
	transactions := h.parser.GetTransactions(address)
	json.NewEncoder(w).Encode(transactions)
}

// SetupRoutes sets up the HTTP routes
func (h *Handler) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/block", h.GetCurrentBlock)
	mux.HandleFunc("/subscribe", h.Subscribe)
	mux.HandleFunc("/transactions", h.GetTransactions)

	return mux
}
