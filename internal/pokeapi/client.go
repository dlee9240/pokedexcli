package pokeapi

import (
	"net/http"
	"time"

	"github.com/dlee9240/pokedexcli/internal/pokecache"
)

type Client struct {
	//create a struct member that is an http.Client
	cache      pokecache.Cache
	httpClient http.Client
}

// New Client
// updated NewClient arguments
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
