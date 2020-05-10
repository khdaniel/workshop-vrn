package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workshop/internal/api"
)

const getJokeAPI = "/api?format=json"

// JokeClient represents Joke client
type JokeClient struct {
	url string
}

// NewJokeClient returns new joke client
func NewJokeClient(baseURL string) *JokeClient {
	return &JokeClient{
		url: baseURL,
	}
}

// GetJoke returns a joke
func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokeAPI

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
