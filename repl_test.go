package main

import (
	"testing"
	"time"

	"github.com/RazafimanantsoaJohnson/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    johnson LEGRAND",
			expected: []string{"johnson", "legrand"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The result of clean input: \n %v \nis not the same as the one we expected: \n %v", actual, c.expected)
			return
		}

		for i, k := range c.expected {
			if actual[i] != k {
				t.Errorf("The result of clean input: \n %v \nis not the same as the one we expected: \n %v", actual, c.expected)
				return
			}
		}
	}
}

func TestExplore(t *testing.T) {
	cases := []struct {
		conf     Config
		expected []string
	}{
		{
			conf: Config{
				LocationBaseUrl:   "https://pokeapi.co/api/v2/location-area/",
				PreviousURL:       "",
				NextURL:           "https://pokeapi.co/api/v2/location-area/",
				CurCommand:        ReceivedCommand{Name: "explore", Params: []string{"pastoria-city-area"}},
				Cache:             pokecache.NewCache(5 * time.Second),
				Pokedex:           make(map[string]Pokemon),
				SupportedCommands: Initializer(),
			},
			expected: []string{
				"tenatacool", "tenatacruel", "magikarp", "gyarados", "remoraid", "octillery", "wingull", "pelipper", "shellos", "gastrodon",
			},
		},
		{
			conf: Config{
				LocationBaseUrl: "https://pokeapi.co/api/v2/location-area/",
				PreviousURL:     "",
				NextURL:         "https://pokeapi.co/api/v2/location-area/",
				CurCommand:      ReceivedCommand{Name: "explore", Params: []string{"canalave-city-area"}},
				Cache:           pokecache.NewCache(5 * time.Second),
			},
			expected: []string{
				"tenatacool", "tenatacruel", "magikarp", "gyarados", "remoraid", "octillery", "wingull", "pelipper", "shellos", "gastrodon",
			},
		},
	}

	// t.Run()
	for _, c := range cases {
		CommandExplore(&c.conf)
	}
}
