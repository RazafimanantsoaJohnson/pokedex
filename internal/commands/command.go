package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

func CommandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(conf *Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n\n")

	for command, value := range *conf.SupportedCommands {
		fmt.Printf("%v: %v\n", command, value.description)
	}
	return nil
}

func CommandMap(conf *Config) error {
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
	if conf.curCommand.name == "mapb" {
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

func CommandExplore(conf *Config) error {
	if len(conf.curCommand.params) <= 0 {
		return fmt.Errorf("this command require 1 argument")
	}
	url := conf.LocationBaseUrl + conf.curCommand.params[0]
	decodeAndShowResult := func(body []byte) error { // first class function to decode and show the result
		var result specificLocationResponse
		conf.cache.Add(url, body)
		err := json.Unmarshal(body, &result)
		if err != nil {
			return err
		}
		for _, encounter := range result.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		return nil
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

func CommandCatch(conf *Config) error {
	pokemonToCatch := conf.curCommand.params[0]
	var body []byte
	url := conf.PokeApiBaseUrl + "pokemon/" + pokemonToCatch
	cachedResult, isPresent := conf.cache.Get(url)
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonToCatch)
	if isPresent {
		return decodeAndCatchPokemon(cachedResult, conf)
	} else {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("a network error occured when trying to get pokemon data")
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("error occured when converting the body in []byte")
		}
		conf.cache.Add(url, body)
		return decodeAndCatchPokemon(body, conf)
	}
}

func decodeAndCatchPokemon(resBody []byte, conf *Config) error {
	var pokemon Pokemon
	err := json.Unmarshal(resBody, &pokemon)
	if err != nil {
		return err
	}
	/*
		Assuming our catching chances are 80-20 for the ranges of base_exp [50-300] and the decrease in chance is linear
		y= mx +b   (where m(rate) and b the y0 'value of y when x=0))
		for us:  y= -0.24x + 92
	*/
	catchingChances := -0.24*float64(pokemon.BaseExperience) + 92 //chances of catching the pokemon based on its base_xp
	randomInt := rand.Intn(100)                                   // a random int to evaluate if we catch it or not
	// fmt.Println(catchingChances)
	if randomInt > int(catchingChances) {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}
	conf.pokedex[pokemon.Name] = pokemon
	fmt.Printf("%v was caught!\n", pokemon.Name)
	return nil
}

func Initializer() *map[string]CliCommand {
	SupportedCommands := map[string]CliCommand{
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
		"explore": {
			name:        "explore",
			description: "List the pokemons we can find in the location given as argument",
			callback:    CommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Capture the pokemon and register it in our pokedex",
			callback:    CommandCatch,
		},
	}
	return &SupportedCommands
}
