package commands

import "fmt"

func commandHelp(config *config) error{
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
	fmt.Println("map: Displays the list of 20 Locations were Pokemons are found")
	fmt.Println("mapb: Displays the list of 20 Previous Locations were Pokemons are found\nexit: Exit the Pokedex")
	// os.Exit(0)
	return  nil
}