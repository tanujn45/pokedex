package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex: ")
	for _, val := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", val.Name)
	}
	fmt.Println()
	return nil
}
