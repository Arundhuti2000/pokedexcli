package main

import "fmt"

func commandHelp(config *config) error{
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
	// os.Exit(0)
	return  nil
}