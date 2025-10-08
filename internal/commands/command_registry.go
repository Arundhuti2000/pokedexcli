package commands

import (
	"time"

	"github.com/Arundhuti2000/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}

type config struct {
	Next      string
	Previous  string
	PokeCache pokecache.Cache
}
type Location_Areas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
}

var cfg = config{
	PokeCache: *pokecache.NewCache(5 * time.Second),
}

// const url = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

// package-level configuration instance used by commands

var Mapcommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex (Usage: exit)",
		Callback:    func() error { return commandExit(&cfg) },
	},
	"help": {
		name:        "help",
		description: "Displays a help message (Usage: help)",
		Callback:    func() error { return commandHelp(&cfg) },
	},
	"map": {
		name:        "map",
		description: "Displays next 20 location areas (Usage: map)",
		Callback:    func() error { return commandMap(&cfg) },
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous 20 location areas (Usage: mapb)",
		Callback:    func() error { return commandMapb(&cfg) },
	},
	"explore": {
		name: "explore",
		description: "list of all the Pokémon located in an area",
		Callback: func() error { return commandExplore(&cfg) },
	}
}