package main

import "github.com/RazafimanantsoaJohnson/pokedexcli/internal/pokecache"

type config struct {
	NextURL     string
	PreviousURL string
	curCommand  string
	cache       pokecache.Cache
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
