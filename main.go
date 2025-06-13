package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scannr := bufio.NewScanner(os.Stdin)
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
	fmt.Print("Pokedex >")
	for scannr.Scan() {
		inputs := CleanInput(scannr.Text())
		if len(inputs) <= 0 {
			fmt.Println("Please provide a command")
			continue
		}
		command, ok := supportedCommands[inputs[0]]
		if ok {
			err := supportedCommands[command.name].callback()
			fmt.Errorf(err.Error())
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Print("Pokedex >")
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CleanInput(text string) []string {
	splittedText := strings.Split(text, " ")
	result := []string{}
	for _, word := range splittedText {
		if word != "" {
			result = append(result, strings.ToLower(word))
		}
	}
	return result
}
