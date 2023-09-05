package handlers

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gophermart/cmd/models"
	"gophermart/cmd/repo"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserLogin string
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	user, modelErr := models.NewUser(r.Body)
	defer r.Body.Close()
	if modelErr != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	_, repoErr := h.repo.CreateUser(user.Login, user.Password)
	if errors.Is(repoErr, repo.ErrUserExists) {
		http.Error(w, "Conflict: User already exists", http.StatusConflict)
		return
	} else if repoErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	token, jwtErr := h.buildJWTString(user.Login)
	if jwtErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
	log.Println("Registered user:", user.Login)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	user, modelErr := models.NewUser(r.Body)
	defer r.Body.Close()
	if modelErr != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	ok, _ := h.repo.CheckPassword(user.Login, user.Password)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, jwtErr := h.buildJWTString(user.Login)
	if jwtErr != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
	log.Println("Logged in user:", user.Login)
}

func (h *Handlers) buildJWTString(userLogin string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(h.conf.TokenExp)),
		},
		UserLogin: userLogin,
	})
	return token.SignedString([]byte(h.conf.SecretKey))
}

func (h *Handlers) getUserLogin(tokenString string) string {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(h.conf.SecretKey), nil
		})
	if err != nil {
		return ""
	}

	if !token.Valid {
		return ""
	}

	return claims.UserLogin
}
