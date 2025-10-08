package commands

import (
	"fmt"

	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)

func CommandHelp(config *registry.Config) error{
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
	fmt.Println("map: Displays the list of 20 Locations were Pokemons are found")
	fmt.Println("mapb: Displays the list of 20 Previous Locations were Pokemons are found")
	fmt.Println("explore: Displays the list of Pokemons present in the location, Usage: explore <location-area-name>\nexit: Exit the Pokedex")
	// os.Exit(0)
	return  nil
}