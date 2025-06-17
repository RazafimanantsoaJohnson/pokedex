package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

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
	if randomInt > int(catchingChances) {
		fmt.Printf("%v escaped!\n", pokemon.Name)
		return nil
	}
	conf.Pokedex[pokemon.Name] = pokemon
	fmt.Printf("%v was caught!\n", pokemon.Name)
	return nil
}

func Initializer() *map[string]CliCommand {
	SupportedCommands := map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "shows 20 locations in our world",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "shows previous 20 locations in our world",
			Callback:    CommandMap,
		},
		"explore": {
			Name:        "explore",
			Description: "List the pokemons we can find in the location given as argument",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Capture the pokemon and register it in our pokedex",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "View details about the pokemon in the Pokedex",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Shows a list of the pokemon the user caught",
			Callback:    CommandPokedex,
		},
	}
	return &SupportedCommands
}
