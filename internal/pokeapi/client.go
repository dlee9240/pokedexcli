package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	//create a struct member that is an http.Client
	httpClient http.Client
}

// New Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
