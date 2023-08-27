package handlers

import (
	"encoding/json"
	"fmt"
	"gophermart/cmd/handlers/models"
	"log"
	"net/http"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	h.repo.CreateUser(user.Username, user.Password)

	log.Println("Registered user:", user.Username)
	w.WriteHeader(http.StatusCreated)
	log.Println(w, "Registration successful")
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}
