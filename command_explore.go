package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a location name")
	}
	loc, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}
	for _, enc := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	fmt.Println()
	return nil
}
