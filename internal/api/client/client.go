package api

import apiModels "goserver/internal/api/models"

type Client interface {
	// GetJoke returns joke
	GetJoke() (*apiModels.JokeResponse, error)
}
