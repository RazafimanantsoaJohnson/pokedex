package commands

import "github.com/RazafimanantsoaJohnson/pokedexcli/internal/pokecache"

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

type Pokemon struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Stats          struct {
		Effort   int `json:"effort"`
		BaseStat int `json:"base_stat"`
	}
}

type Config struct {
	LocationBaseUrl   string
	PokeApiBaseUrl    string
	NextURL           string
	PreviousURL       string
	curCommand        ReceivedCommand
	cache             pokecache.Cache
	pokedex           map[string]Pokemon
	SupportedCommands *map[string]CliCommand
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(conf *Config) error
}

type ReceivedCommand struct {
	Name   string
	Params []string
}
