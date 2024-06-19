package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex help.")
	fmt.Println("Here are your availabe command:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s:  %s\n", cmd.name, cmd.description)
	}
	return nil
}
