package main

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func GetCommands() map[string]cliCommand {
	commands := map[string]cliCommand{}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Get the next page of locations",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Get the previous page of locations",
		callback:    commandMapb,
	}
	commands["explore"] = cliCommand{
		name:        "explore <location_name>",
		description: "Explore a location",
		callback:    commandExplore,
	}
	commands["catch"] = cliCommand{
		name:        "catch <pokemon_name>",
		description: "Catch a pokemon",
		callback:    commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect <pokemon_name>",
		description: "Inspect a caught pokemon",
		callback:    commandInspect,
	}
	return commands
}

func commandInvalid() {
	fmt.Println("Invalid command!")
}
