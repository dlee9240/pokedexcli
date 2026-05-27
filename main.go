package main

import (
	"time"

	"github.com/dlee9240/pokedexcli/internal/pokeapi"
)

func main() {
	//need to set up the pokeapi client etc

	pokeClient := pokeapi.NewClient(5 * time.Second)

	//setting up the actual config and assigning the pokeClient above to the struct...
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
