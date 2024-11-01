package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}
	fmt.Print("\n")
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, result := range locationResp.Results {
		fmt.Println(result.Name)
	}
	fmt.Print("\n")
	return nil
}
