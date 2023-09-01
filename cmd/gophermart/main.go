package main

import (
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
	h := handlers.New(repo, conf)

	log.Println("Starting server on address", conf.RunAddress)
	router := h.NewRouter()
	http.ListenAndServe(conf.RunAddress, router)
}
