package handlers

import (
	"encoding/json"
	"gophermart/cmd/handlers/models"
	"io"
	"net/http"
)

func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userLogin := h.getUserLogin(token[7:])
	if userLogin == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, userErr := h.repo.GetUser(userLogin)
	if userErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	orderNumber, _ := io.ReadAll(r.Body)
	// TODO: add validation
	_, repoErr := h.repo.CreateOrder(string(orderNumber), userID)
	if repoErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h *Handlers) GetOrders(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userLogin := h.getUserLogin(token[7:])
	if userLogin == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, userErr := h.repo.GetUser(userLogin)
	if userErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	orders, repoErr := h.repo.GetOrders(userID)
	if repoErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var jsonOrders []models.Order
	for _, order := range orders {
		jsonOrders = append(jsonOrders, models.Order{
			Number:    order.Number,
			CreatedAt: order.CreatedAt,
			Status:    order.Status,
			Accrual:   order.Accrual,
		})
	}
	jsonData, err := json.Marshal(jsonOrders)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
