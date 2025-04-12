package main

import (
	"log"
	"net/http"

	"go-rest-api/internal/api"
	"go-rest-api/internal/handlers"
	"go-rest-api/internal/services"
)

func main() {
	service := services.NewItemService()
	handler := handlers.NewItemHandler(service)
	router := api.NewRouter(handler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
