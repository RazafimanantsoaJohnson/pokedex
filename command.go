package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var SupportedCommands map[string]cliCommand

func CommandExit(conf *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(conf *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n\n")

	for command, value := range SupportedCommands {
		fmt.Printf("%v: %v\n", command, value.description)
	}
	return nil
}

func CommandMap(conf *config) error {
	url := conf.NextURL
	decodeAndShowResult := func(body []byte) error { // first class function to decode and show the result
		var result locationResponse
		conf.cache.Add(url, body)
		err := json.Unmarshal(body, &result)
		if err != nil {
			return err
		}

		conf.NextURL = result.Next
		conf.PreviousURL = result.Previous
		for _, location := range result.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
	if conf.curCommand == "mapb" {
		url = conf.PreviousURL
	}
	var body []byte
	cachedResult, isPresent := conf.cache.Get(url)
	if isPresent {
		// fmt.Println("We are using cache here BTW :3")
		return decodeAndShowResult(cachedResult)
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("An error happened when converting the body in []byte")
		}
		return decodeAndShowResult(body)
	}
}

func Initializer() {
	SupportedCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "shows 20 locations in our world",
			callback:    CommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "shows previous 20 locations in our world",
			callback:    CommandMap,
		},
	}
}
