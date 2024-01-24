package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/num30/config"

	handler "goserver/internal"
	"goserver/internal/api/jokes"
	appConf "goserver/internal/config"
)

func main() {
	cfg := appConf.Server{}

	cfgReader := config.NewConfReader("config").WithSearchDirs("./", ".")
	err := cfgReader.Read(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", cfg)

	apiClient := jokes.NewJokeClient(cfg.JokeUrl)

	envErr := os.Setenv("PORT", cfg.Port)
	if envErr != nil {
		log.Fatal(envErr)
	}

	var port = envPortOr("3000")
	h := handler.NewHandler(apiClient)
	r := chi.NewRouter()

	r.Get("/", h.Joke)
	log.Println("Starting Server")
	error := http.ListenAndServe(port, r)
	log.Fatal(error)
	log.Println("Shutting server down")
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
