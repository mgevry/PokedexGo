package pokeapi

import (
	"net/http"
	"time"
)
//creates a custom client with a timeout, used for pokeapi client
type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}