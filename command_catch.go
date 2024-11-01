package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("you must provide a pokemon name")
	}
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	res := rand.IntN(pokemonResp.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n\n", pokemonResp.Name)

	cfg.caughtPokemon[args[0]] = pokemonResp
	return nil
}
