package main

import (
	"gophermart/cmd/config"
	"gophermart/cmd/handlers"
	"gophermart/cmd/repo/postgres"
	"log"
	"net/http"
)

func main() {
	log.Println("Initializing server configuration and handlers")
	conf := config.NewConfig()
	repo := postgres.NewPostgresRepo(conf)
	h := handlers.New(repo, conf)

	log.Println("Starting server on address", conf.RunAddress)
	router := h.NewRouter()
	http.ListenAndServe(conf.RunAddress, router)
}
