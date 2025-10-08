package registry

import "github.com/Arundhuti2000/pokedexcli/internal/pokecache"

type Config struct {
	Next      string
	Previous  string
	PokeCache pokecache.Cache
}