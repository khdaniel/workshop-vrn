package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workshop/internal/api"
)

const getJokeAPI = "/api?format=json"

type JokeClientAPI struct {
	url string
}

func (jc *JokeClientAPI) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokeAPI

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %v", err)
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
