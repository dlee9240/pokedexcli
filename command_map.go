package main

import (
	"errors"
	"fmt"
)

func commandMapf(configuration *config) error {
	//get the locations response...
	locationsResp, err := configuration.pokeapiClient.ListLocations(configuration.nextLocationsURL)
	if err != nil {
		return err
	}

	configuration.nextLocationsURL = locationsResp.Next
	configuration.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(configuration *config) error {
	//do the first check if this is the first page
	if configuration.previousLocationsURL == nil {
		return errors.New("You're on the first page")
	}

	//get the location response...
	locationsResp, err := configuration.pokeapiClient.ListLocations(configuration.previousLocationsURL)
	if err != nil {
		return err
	}

	configuration.nextLocationsURL = locationsResp.Next
	configuration.previousLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
