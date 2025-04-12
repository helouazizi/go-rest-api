package main

import (
	"fmt"
	"log"
	"net/http"

	"go-rest-api/internal/api"
	"go-rest-api/internal/dependencies"
	"go-rest-api/pkg/logger"
)

func main() {
	// Get the current working directory
	logger, err := logger.Create_Logger()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer logger.Close()

	deps := dependencies.NewDependencies()
	router := api.NewRouter(deps)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
