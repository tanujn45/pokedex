package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	return commands
}

func commandInvalid() {
	fmt.Println("Invalid command!")
}

func commandHelp() error {
	fmt.Print("\nWelcome to Pokedex!\nUsage:\n\n")
	for _, value := range GetCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	fmt.Print("\n")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
