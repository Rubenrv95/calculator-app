package main

import (
	"log"
	"net/http"

	"github.com/Rubenrv95/calculator-app/backend/internal/handlers"
)

func main() {
	router := handlers.NewRouter()

	addr := ":8080"
	log.Printf("calculator backend listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
