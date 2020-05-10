package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

type Handler struct {
	JokeClient api.Client
	customJoke string
}

func NewHandler(jokeClient api.Client, customJoke string) *Handler {
	return &Handler{
		JokeClient: jokeClient,
		customJoke: customJoke,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	if h.customJoke != "" {
		fmt.Fprint(w, h.customJoke)
		return
	}

	joke, err := h.JokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, joke.Joke)
}
