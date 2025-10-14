# Pokedex CLI

A simple command-line Pokedex I built while learning Go. It lets you explore the Pokemon world, catch Pokemon, and build your own collection using the PokeAPI.

## What it does

- Browse through location areas to find Pokemon
- Explore specific areas to see what Pokemon are there
- Try catching Pokemon (with some luck involved!)
- Keep track of everything you catch
- Check out detailed stats for your Pokemon
- Caches API responses so it runs faster

## Getting Started

You'll need Go 1.25.0+ installed...

```bash
git clone https://github.com/Arundhuti2000/pokedexcli.git
cd pokedexcli
go build
./pokedexcli
```

That's it!

## How to use it

Type commands at the `Pokedex >` prompt:

**`help`** - Shows all commands  
**`map`** - Next 20 locations  
**`mapb`** - Previous 20 locations  
**`explore <location>`** - See Pokemon in that area  
**`catch <pokemon>`** - Try to catch it  
**`inspect <pokemon>`** - View stats (only works if you caught it)  
**`pokedex`** - See everything you've caught  
**`exit`** - Quit  

### Quick example

```
Pokedex > help
Welcome to the Pokedex!
Usage:
help: Displays a help message
map: Displays the list of 20 Locations were Pokemons are found
...

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore canalave-city-area
Number of pokemon encounters: 5
tentacool
tentacruel
staryu
...

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!

Pokedex > pokedex
Your Pokedex:
  - pikachu

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  - hp: 35
  - attack: 55
  - defense: 40
  ...
Types:
  - electric

Pokedex > exit
Closing the Pokedex... Goodbye!
```

## Project structure

```
pokedexcli/
├── main.go                      # Where it all starts
├── internal/
│   ├── client/
│   │   └── client.go           # Command registry
│   ├── commands/               # All the commands
│   │   ├── command_cleanInput.go
│   │   ├── command_help.go
│   │   ├── command_exit.go
│   │   ├── command_map.go
│   │   ├── command_explore.go
│   │   ├── command_catch.go
│   │   ├── command_inspect.go
│   │   └── command_pokedex.go
│   ├── pokecache/              # Cache implementation
│   │   ├── pokecache.go
│   │   └── pokecache_test.go
│   └── registry/               # Data structures
│       ├── cliCommand.go
│       ├── config.go
│       ├── location-areas.go
│       └── pokemon.go
└── repl_test.go
```

## How it works

### The cache
I added a simple cache to avoid hammering the PokeAPI. Responses are cached for 10 minutes and cleaned up automatically in the background. Uses mutexes to keep things thread-safe.

### Catching Pokemon
The catch probability is based on each Pokemon's base experience - basically, stronger Pokemon are harder to catch. There's always at least a 5% chance of success though, so keep trying!

## Running tests

```bash
go test ./...
```

Or test specific packages:
```bash
go test ./internal/pokecache
```

## Data source

All Pokemon data comes from [PokeAPI v2](https://pokeapi.co/docs/v2).

## Want to contribute?

Feel free to fork and submit PRs. This is a learning project so all contributions are welcome!

## Credits

Thanks to PokeAPI for the free API - this wouldn't exist without it.

---

Built this while learning Go. Still learning, so the code might not be perfect!
