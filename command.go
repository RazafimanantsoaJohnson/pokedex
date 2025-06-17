package main

import (
	"encoding/json"
	"fmt"
	"io"
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
		fmt.Printf("%v: %v\n", command, value.Description)
	}
	return nil
}

func CommandMap(conf *Config) error {
	url := conf.NextURL
	decodeAndShowResult := func(body []byte) error { // first class function to decode and show the result
		var result locationResponse
		conf.Cache.Add(url, body)
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
	if conf.CurCommand.Name == "mapb" {
		url = conf.PreviousURL
	}
	var body []byte
	cachedResult, isPresent := conf.Cache.Get(url)
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
	if len(conf.CurCommand.Params) <= 0 {
		return fmt.Errorf("this command require 1 argument")
	}
	url := conf.LocationBaseUrl + conf.CurCommand.Params[0]
	decodeAndShowResult := func(body []byte) error { // first class function to decode and show the result
		var result specificLocationResponse
		conf.Cache.Add(url, body)
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
	cachedResult, isPresent := conf.Cache.Get(url)
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
	pokemonToCatch := conf.CurCommand.Params[0]
	var body []byte
	url := conf.PokeApiBaseUrl + "pokemon/" + pokemonToCatch
	cachedResult, isPresent := conf.Cache.Get(url)
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
		conf.Cache.Add(url, body)
		return decodeAndCatchPokemon(body, conf)
	}
}

func CommandInspect(conf *Config) error {
	pokemonToInspect := conf.CurCommand.Params[0]
	pokemon, ok := conf.Pokedex[pokemonToInspect]
	if !ok {
		fmt.Printf("You did not catch %v yet\n", pokemonToInspect)
		return nil
	}
	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%v:%v\n", stat.BaseStat, stat.Stat.Name)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("\t-%v\n", t.Type.Name)
	}
	return nil
}

func CommandPokedex(conf *Config) error {
	fmt.Printf("Your Pokedex:\n")
	for _, pokemon := range conf.Pokedex {
		fmt.Printf("\t- %v\n", pokemon.Name)
	}
	return nil
}
