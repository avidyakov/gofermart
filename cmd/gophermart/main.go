package main

import (
	"gophermart/cmd/loyalty"
	"log"
	"net/http"
)

import (
	"gophermart/cmd/config"
	"gophermart/cmd/handlers"
	"gophermart/cmd/repo/postgres"
)

func main() {
	log.Println("Initializing server configuration and handlers")
	conf := config.NewConfig()
	repo := postgres.NewRepo(conf)
	accrual := loyalty.NewAccrualSystem(conf.AccrualAddress)
	h := handlers.New(repo, conf, accrual)

	log.Println("Starting server on address", conf.RunAddress)
	router := h.NewRouter()
	http.ListenAndServe(conf.RunAddress, router)
}
