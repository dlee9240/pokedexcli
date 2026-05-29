package main

import (
	"time"

	"github.com/dlee9240/pokedexcli/internal/pokeapi"
)

func main() {
	//need to set up the pokeapi client etc

	//added the second argument of 5 minutes
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)

	//setting up the actual config and assigning the pokeClient above to the struct...
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
