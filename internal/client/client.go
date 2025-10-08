package client

import (
	"time"

	"github.com/Arundhuti2000/pokedexcli/internal/commands"
	"github.com/Arundhuti2000/pokedexcli/internal/pokecache"
	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)
var cfg = registry.Config{
	PokeCache: *pokecache.NewCache(5 * time.Second),
}

// const url = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

// package-level configuration instance used by commands

var Mapcommands = map[string]registry.CliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex (Usage: exit)",
		Callback:    func() error { return commands.CommandExit(&cfg) },
	},
	"help": {
		Name:        "help",
		Description: "Displays a help message (Usage: help)",
		Callback:    func() error { return commands.CommandHelp(&cfg) },
	},
	"map": {
		Name:        "map",
		Description: "Displays next 20 location areas (Usage: map)",
		Callback:    func() error { return commands.CommandMap(&cfg) },
	},
	"mapb": {
		Name:        "mapb",
		Description: "Displays previous 20 location areas (Usage: mapb)",
		Callback:    func() error { return commands.CommandMapb(&cfg) },
	},
	"explore": {
		Name: "explore",
		Description: "list of all the Pokémon located in an area",
		CallbackWithArg: func(location string) error { return commands.CommandExplore(&cfg, location) },
	},
}