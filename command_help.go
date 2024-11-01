package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Print("\nWelcome to Pokedex!\nUsage:\n\n")
	for _, value := range GetCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Print("\n")
	return nil
}
