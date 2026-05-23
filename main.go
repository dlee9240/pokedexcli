package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

map[string]cliCommand{
	"exit": {
		name: "exit",
		description: "Exits the Pokedex",
		callback: commandExit,
		},
}



func main() {
	//remove hello world logic
	//fmt.Println("Hello, World!")
	//scanner := bufio.NewScanner(os.Stdin)


	//removed the prints the first word back..
}
