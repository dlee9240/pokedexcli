package main

import (
	"errors"
	"fmt"
)

func commandInspect(configuration *config, args ...string) error {
	//do something
	if len(args) != 1 {
		fmt.Println("You must specify a pokemon name")
	}

	name := args[0]

	pokemon, ok := configuration.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Println("  -", pType.Type.Name)
	}
	return nil

}
