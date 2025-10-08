package commands

import "fmt"

func commandExplore(config *config, location string) error{
	fmt.Println("explore: Displays the list of pokemon found in that area\tUsage: explore <location-name>")
	url:= config.Next
	if url == ""{
		url = "https://pokeapi.co/api/v2/location-area/"+ location +"?limit=20"
	}
	// os.Exit(0)
	return  nil
}