package pokeapi

import (
	"net/http"
	"time"
	"github.com/mgevry/pokedex/internal"
)
//creates a custom client with a timeout, used for pokeapi client
type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}