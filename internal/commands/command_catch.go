package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)

func CommandCatch(config *registry.Config, name string) error{
	// fmt.Println("explore: Displays the list of pokemon found in that area")

	url := "https://pokeapi.co/api/v2/location-area/"+ name 
	var body []byte
	pokecache, ok:= config.PokeCache.Get(url)
	
	if !ok{
		fmt.Println("Cache miss - fetching from API")
		req, err:= http.Get(url)
		if err != nil{
			fmt.Println("Error fetching:", err)
			return err
		}
		body, err = io.ReadAll(req.Body)
		if err!=nil{
			fmt.Println("Error reading body:", err)
			return err
		}
		defer req.Body.Close()
		config.PokeCache.Add(url, body)
		// fmt.Println(config)
		if req.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", req.StatusCode, body)
		}
		var pokemon_by_location_area registry.Pokemon_by_location_area
		err = json.Unmarshal(body, &pokemon_by_location_area)
		if err!=nil{
			fmt.Println("Error unmarshaling:", err)
			return err
		}
		fmt.Println("Number of pokemon encounters:", len(pokemon_by_location_area.Pokemon_encounters))
		for _, pokemon_encounter := range pokemon_by_location_area.Pokemon_encounters {
			fmt.Println(pokemon_encounter.Pokemon.Name)
		}
	} else{
		fmt.Println("Cache hit")
		var pokemon_by_location_area registry.Pokemon_by_location_area
		err := json.Unmarshal(pokecache, &pokemon_by_location_area)
		if err!=nil{
			fmt.Println("Error unmarshaling from cache:", err)
			return err
		}
		fmt.Println("Number of pokemon encounters:", len(pokemon_by_location_area.Pokemon_encounters))
		for _, pokemon_encounter := range pokemon_by_location_area.Pokemon_encounters {
			fmt.Println(pokemon_encounter.Pokemon.Name)
		}
	}
	return  nil
}