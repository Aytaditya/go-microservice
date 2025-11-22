package main

import (
	"fmt"
	"net/http"

	service "github.com/Aytaditya/go-microservice/internal/http"
)

func main() {
	// config file load
	// database connection
	router := http.NewServeMux()
	router.HandleFunc("GET /", service.GiveService())

	server := http.Server{
		Addr:    ":3001",
		Handler: router,
	}
	fmt.Println("Service 1 is running on port 3001")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
