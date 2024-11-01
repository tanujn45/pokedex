package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you need to provide a pokemon name")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		fmt.Println()
		return nil
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, val := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types: ")
	for _, val := range pokemon.Types {
		fmt.Printf("  - %s\n", val.Type.Name)
	}
	fmt.Println()
	return nil
}
