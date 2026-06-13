package main

import (
	"fmt"
)

func commandPokedex(configuration *config, args ...string) error {
	//takes no arguments...  The args argument can be left black so no need to check the length etc.

	fmt.Println("Your Pokedex:")
	for _, pokemon := range configuration.caughtPokemon {
		fmt.Printf("  -%v\n", pokemon.Name)
	}

	return nil
}
