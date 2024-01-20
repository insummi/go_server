package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	handler "goserver/internal"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
