package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	handler "goserver/internal"
	config "goserver/internal/config"
)

func main() {
	cfg := config.Server{}
	log.Print(cfg)
	var port = envPortOr("3000")
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/", h.Hello)
	log.Print("Starting Server")
	err := http.ListenAndServe(port, r)
	log.Fatal(err)
	log.Print("Shutting server down")
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
