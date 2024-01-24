package handler

import (
	"fmt"
	api "goserver/internal/api/client"
	"net/http"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient,
	}
}

func (h *Handler) Joke(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, joke.Joke)
}
