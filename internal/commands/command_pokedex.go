package commands

import (
	"fmt"

	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)

func CommandPokedex(config *registry.Config) error {
	fmt.Println("Your Pokedex:")
	if len(config.CaughtPokemon) == 0 {
		fmt.Println("  You haven't caught any Pokemon yet!")
	}else {
		for name := range config.CaughtPokemon {
			fmt.Printf("  - %s\n", name)
		}
	}
	return nil
}