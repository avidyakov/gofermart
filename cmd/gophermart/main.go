package main

import (
	"gophermart/cmd/config"
	"gophermart/cmd/handlers/routers"
	"log"
	"net/http"
)

func main() {
	log.Println("Initializing server configuration and handlers")
	config := config.NewConfig()

	log.Println("Starting server on address", config.RunAddress)
	router := routers.NewRouter()
	http.ListenAndServe(config.RunAddress, router)
}
