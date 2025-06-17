package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RazafimanantsoaJohnson/pokedexcli/internal/pokecache"
)

func main() {
	scannr := bufio.NewScanner(os.Stdin)
	var conf Config = Config{
		LocationBaseUrl:   "https://pokeapi.co/api/v2/location-area/",
		PokeApiBaseUrl:    "https://pokeapi.co/api/v2/",
		PreviousURL:       "",
		NextURL:           "https://pokeapi.co/api/v2/location-area/",
		SupportedCommands: Initializer(),
		CurCommand:        ReceivedCommand{Name: "", Params: []string{}},
		Cache:             pokecache.NewCache(5 * time.Second),
		Pokedex:           make(map[string]Pokemon),
	}

	fmt.Print("Pokedex >")
	for scannr.Scan() {
		inputs := CleanInput(scannr.Text())
		if len(inputs) <= 0 {
			fmt.Println("Please provide a command")
			fmt.Print("Pokedex >")
			continue
		}
		command, ok := (*conf.SupportedCommands)[inputs[0]]
		if ok {
			conf.CurCommand.Name = command.Name
			if len(inputs) > 1 {
				conf.CurCommand.Params = inputs[1:]
			}
			err := (*conf.SupportedCommands)[command.Name].Callback(&conf)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Print("Pokedex >")
	}
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
