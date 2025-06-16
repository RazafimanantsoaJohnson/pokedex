package main

import "github.com/RazafimanantsoaJohnson/pokedexcli/internal/pokecache"

type config struct {
	LocationBaseUrl string
	NextURL         string
	PreviousURL     string
	curCommand      receivedCommand
	cache           pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config) error
}

type locationResponse struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []locationResults `json:"results"`
}

type locationResults struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type specificLocationResponse struct {
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon pokemonInLocation `json:"pokemon"`
}

type pokemonInLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type receivedCommand struct {
	name   string
	params []string
}
