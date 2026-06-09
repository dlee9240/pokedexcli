package main

import (
	"errors"
	"fmt"
)

func commandExplore(configuration *config, args ...string) error {
	//do something
	//display a list of all Pokemon located at the location above...
	//changed the args ...string from locations string

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := configuration.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Printf("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
