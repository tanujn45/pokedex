package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := GetCommands()

	for {
		fmt.Print("Pokedex > ")
		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}
		inputSlice := cleanInput(input)
		command, ok := commands[inputSlice[0]]
		if !ok {
			commandInvalid()
			continue
		}
		command.callback()
	}
}

func cleanInput(input string) []string {
	output := strings.ToLower(input)
	words := strings.Fields(output)
	return words
}
