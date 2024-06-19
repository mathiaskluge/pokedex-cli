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

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		cmd, ok := commands[input]
		if !ok {
			fmt.Println("Command not found. Type 'help' for a list of available commands.")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
		}

		if input == "exit" {
			break
		}
	}
}

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting the Pokedex. Goodbye!")
	return nil
}
