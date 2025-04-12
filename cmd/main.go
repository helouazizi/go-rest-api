package main

import (
	"log"
	"net/http"

	"go-rest-api/internal/api"
	"go-rest-api/internal/dependencies"
)

func main() {
	deps := dependencies.NewDependencies()
	router := api.NewRouter(deps)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
