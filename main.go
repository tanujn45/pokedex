package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tanujn45/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(5 & time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}
		inputSlice := cleanInput(input)
		command, ok := GetCommands()[inputSlice[0]]
		if !ok {
			commandInvalid()
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
            fmt.Println()
		}
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
