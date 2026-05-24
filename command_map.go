package main

import (
	"io"
	"log"
	"net/http"
)

func commandMap() error {
	//add the GET Request here and map it to the Struct
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}

	//take the body and marshal it into a GO STRUCT which should be of type CLICommand?
	//gonna stop here for now
	
	return nil
}
