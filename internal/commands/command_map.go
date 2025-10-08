package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(config *config) error{
	fmt.Println("Displays all the locations where Pokemons are found")
	// fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
	// os.Exit(0)
	url:= config.Next
	if url == ""{
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
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
		var location_areas Location_Areas
		err = json.Unmarshal(body, &location_areas)
		if err!=nil{
			return err
		}
		for _, area := range location_areas.Results {
			fmt.Println(area.Name)
		}

		config.Next = location_areas.Next
		config.Previous = location_areas.Previous
	}else{
		var location_areas Location_Areas
		err := json.Unmarshal(pokecache, &location_areas)
		if err != nil {
			return err
		}
		for _, area := range location_areas.Results {
			fmt.Println(area.Name)
		}

		config.Next = location_areas.Next
		config.Previous = location_areas.Previous
	}
	return  nil
}

func commandMapb(config *config) error{
	fmt.Println("Displays the previous 20 locations where Pokemons are found")
	// fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
	// os.Exit(0)
	url:= config.Previous
	if url == ""{
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
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
		var location_areas Location_Areas
		err = json.Unmarshal(body, &location_areas)
		if err!=nil{
			return err
		}
		for _, area := range location_areas.Results {
			fmt.Println(area.Name)
		}

		config.Next = location_areas.Next
		config.Previous = location_areas.Previous
	} else{
		var location_areas Location_Areas
		err := json.Unmarshal(pokecache, &location_areas)
		if err != nil {
			return err
		}
		for _, area := range location_areas.Results {
			fmt.Println(area.Name)
		}

		config.Next = location_areas.Next
		config.Previous = location_areas.Previous
	}
	
	return  nil
}