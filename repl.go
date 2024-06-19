package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()

	fmt.Println("This will echo everything you type!")
	for {
		fmt.Print(" > ")

		scanner.Scan()
		input := cleanInput(scanner.Text())

		// enables empty returns
		if len(input) == 0 {
			continue
		}

		command, ok := availableCommands[input[0]]
		if !ok {
			fmt.Println("Command not found. Type 'help' for a list of available commands.")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("There was a problem executing the command: %s", err)
		}
	}
}

func cleanInput(str string) []string {
	lowered_words := strings.Fields(strings.ToLower(str))
	return lowered_words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Closes the Pokedex",
			callback:    callbackExit,
		},
		"help": {
			name:        "help",
			description: "Shows all available commands and what they do",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Shows 20 locations in the Pokemon world",
			callback:    callbackMap,
		},
	}
}
