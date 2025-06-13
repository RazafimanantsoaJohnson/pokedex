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

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(listOfCommands map[string]cliCommand) func() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n\n")

	return func() error {
		for command, value := range listOfCommands {
			fmt.Printf("%v: %v", command, value.description)
		}
		return nil
	}
}
