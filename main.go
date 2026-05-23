package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//remove hello world logic
	//fmt.Println("Hello, World!")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		//strings.ToLower(text)
		result := strings.SplitN(strings.ToLower(text), " ", 2)
		fmt.Printf("Your command was: %s\n", result[0])
	}
}
