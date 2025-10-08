package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)

func CommandExplore(config *registry.Config, location string) error{
	fmt.Println("explore: Displays the list of pokemon found in that area\tUsage: explore <location-area-name>")
	// url:= config.Next
	// if url == ""{
		
	// }
	url := "https://pokeapi.co/api/v2/location-area/"+ location +"?limit=20"
	pokecache, ok:= config.PokeCache.Get(url)
	if !ok{
		req, err:= http.Get(url)
		if err != nil{
			return err
		}
		body, err := io.ReadAll(req.Body)
		if err!=nil{
			return err
		}
		defer req.Body.Close()
		config.PokeCache.Add(url, body)
		if req.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", req.StatusCode, body)
		}
		var pokemon_by_location_area registry.Pokemon_by_location_area
		err = json.Unmarshal(body, &pokemon_by_location_area)
		if err!=nil{
			return err
		}
		for _, pokemon_encounter := range pokemon_by_location_area.Pokemon_encounters {
			fmt.Println(pokemon_encounter.Pokemon.Name)
		}
		
	} else{
		var pokemon_by_location_area registry.Pokemon_by_location_area
		err := json.Unmarshal(pokecache, &pokemon_by_location_area)
		if err!=nil{
			return err
		}
		for _, pokemon_encounter := range pokemon_by_location_area.Pokemon_encounters {
			fmt.Println(pokemon_encounter.Pokemon.Name)
		}
	}
	// os.Exit(0)
	return  nil
}