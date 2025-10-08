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
	url:= config.Next
	if url == ""{
		url = "https://pokeapi.co/api/v2/location-area/"+ location +"?limit=20"
	}
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
		var Pokemon_by_location_area Pokemon_by_location_area
		err = json.Unmarshal(body, &location_areas)
		if err!=nil{
			return err
		}
		for _, area := range location_areas.Results {
			fmt.Println(area.Name)
		}
	}
	// os.Exit(0)
	return  nil
}