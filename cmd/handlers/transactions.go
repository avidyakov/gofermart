package handlers

import (
	"encoding/json"
	"gophermart/cmd/models"
	"log"
	"net/http"
)

func (h *Handlers) GetBalance(w http.ResponseWriter, r *http.Request) {
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

	balance, balanceErr := h.repo.GetBalance(userID)
	if balanceErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	used, repoErr := h.repo.GetUsed(userID)
	if repoErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := models.Balance{
		Current:   balance,
		Withdrawn: used,
	}
	jsonData, err := json.Marshal(response)
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
	log.Printf("User %s get balance", userLogin)
}
