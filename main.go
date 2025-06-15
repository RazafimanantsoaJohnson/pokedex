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
	Initializer()
	var conf config = config{
		PreviousURL: "url1",
		NextURL:     "https://pokeapi.co/api/v2/location-area/",
		curCommand:  "",
		cache:       pokecache.NewCache(5 * time.Second),
	}

	fmt.Print("Pokedex >")
	for scannr.Scan() {
		inputs := CleanInput(scannr.Text())
		if len(inputs) <= 0 {
			fmt.Println("Please provide a command")
			fmt.Print("Pokedex >")
			continue
		}
		command, ok := SupportedCommands[inputs[0]]
		if ok {
			conf.curCommand = command.name
			err := SupportedCommands[command.name].callback(&conf)
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
