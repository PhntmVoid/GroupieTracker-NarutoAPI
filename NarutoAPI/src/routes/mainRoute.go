package routes

import (
	"log"
	"net/http"
)

func InitServer() {
	homeRoute()
	charactersRoute()
	clansRoute()
	kekkeiGenkaiRoute()
	tailedBeastsRoute()
	aboutRoute()

	server := &http.Server{
		Addr: "localhost:8080",
	}

	log.Printf("Starting server at http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
