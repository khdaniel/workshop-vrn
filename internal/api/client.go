package api

// Client interacts with 3rd party joke API
type Client interface {
	// GetJoke returns one joke
	GetJoke() (*JokeResponse, error)
}
