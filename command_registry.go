package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	Next     string
	Previous string
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

// const url = "https://pokeapi.co/api/v2/location-area?offset=20&limit=20"

// package-level configuration instance used by commands
var cfg = config{}

var mapcommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex (Usage: exit)",
		callback:    func() error { return commandExit(&cfg) },
	},
	"help": {
		name:        "help",
		description: "Displays a help message (Usage: help)",
		callback:    func() error { return commandHelp(&cfg) },
	},
	"map": {
		name:        "map",
		description: "Displays next 20 location areas (Usage: map)",
		callback:    func() error { return commandMap(&cfg) },
	},
	"mapb": {
		name:        "mapb",
		description: "Displays previous 20 location areas (Usage: mapb)",
		callback:    func() error { return commandMapb(&cfg) },
	},
}