package main

import (
	"strings"
)

func cleanInput(text string) []string {
	//return an empty string for now
	//return []string{}
	//takes a text string and splits it into a slice of strings.. Remove the whitespace etc
	//var result []string
	result := strings.Fields(strings.ToLower(text))
	return result
}

//this is boot.dev solution
/*
package main

import (
	"strings"
)

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

*/
