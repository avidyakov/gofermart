package handlers

import (
	"encoding/json"
	"errors"
	"gophermart/cmd/handlers/models"
	"gophermart/cmd/repo"
	"log"
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

	orderNumber, validationErr := models.NewOrderInput(r.Body)
	if validationErr != nil {
		http.Error(w, validationErr.Error(), http.StatusUnprocessableEntity)
		return
	}
	_, repoErr := h.repo.CreateOrder(string(*orderNumber), userID)
	if errors.Is(repoErr, repo.ErrOrderExists) {
		http.Error(w, "Order already exists", http.StatusConflict)
		return
	} else if errors.Is(repoErr, repo.ErrOrderAlreadyUploaded) {
		w.WriteHeader(http.StatusOK)
		return
	} else if repoErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	log.Printf("User %s create order %s", userLogin, string(*orderNumber))
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
	var jsonOrders []models.OrderOutput
	for _, order := range orders {
		jsonOrders = append(jsonOrders, models.OrderOutput{
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

	log.Printf("User %s get orders", userLogin)
}
