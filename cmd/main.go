package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	handler "goserver/internal"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/", h.Hello)
	log.Print("Starting Server")
	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
	log.Print("Shutting server down")
}
