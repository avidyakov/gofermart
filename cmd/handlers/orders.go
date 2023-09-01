package handlers

import (
	"io/ioutil"
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

	orderNumber, _ := ioutil.ReadAll(r.Body)
	// TODO: add validation
	_, repoErr := h.repo.CreateOrder(string(orderNumber), userID)
	if repoErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
