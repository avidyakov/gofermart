package handlers

import (
	"errors"
	"fmt"
	"gophermart/cmd/handlers/models"
	"gophermart/cmd/repo"
	"log"
	"net/http"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	user, err := models.NewUser(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = h.repo.CreateUser(user.Login, user.Password)
	if errors.Is(err, repo.ErrUserExists) {
		http.Error(w, "Conflict: User already exists", http.StatusConflict)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Registered user:", user.Login)
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}
