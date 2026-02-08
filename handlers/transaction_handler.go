package handlers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
)

type TransactionHandler struct {
	service services.TransactionServiceInterface
}

func NewTransactionHandler(service services.TransactionServiceInterface) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req models.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Error(http.StatusBadRequest, models.ErrInternal, err.Error()))
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Error(http.StatusInternalServerError, models.ErrInternal, err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.Success(http.StatusCreated, "Transaction created successfully", transaction))
}
