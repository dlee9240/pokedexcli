package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Locations struct {
	Name string
	Url  string
}

// Config struct to store the information from the Location API
type Config struct {
	Count    int
	Next     string
	Previous string
	Results  []Locations
}

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

	//body is the response body from the API Call. Need to marshal it out to a GO Struct..

	//prints the bytes sent over the wire from the API Call.
	//fmt.Println(string(body))

	var response Config
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}
	//check this out to see if the marshal worked as I expect... WORKS!!!
	fmt.Println(response.Count)
	fmt.Println(response.Next)
	fmt.Println(response.Previous)
	return nil
}
