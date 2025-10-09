package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)

func attemptCatch(baseExperience int) bool {
	catchThreshold := 50 + (baseExperience / 10)
	
	if catchThreshold > 95 {
		catchThreshold = 95
	}
	
	randomValue := rand.Intn(100)
	
	return randomValue >= catchThreshold
}

func CommandCatch(config *registry.Config, name string) error{
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	
	url := "https://pokeapi.co/api/v2/pokemon/"+ name 
	
	var body []byte
	pokecache, ok:= config.PokeCache.Get(url)
	
	if !ok{
		req, err:= http.Get(url)
		if err != nil{
			fmt.Println(err)
			return err
		}
		defer req.Body.Close()
		
		if req.StatusCode > 299 {
			fmt.Println(err)
			return fmt.Errorf("Response failed with status code: %d", req.StatusCode)
		}
		
		body, err = io.ReadAll(req.Body)
		if err!=nil{
			fmt.Println(err)
			return err
		}
		
		config.PokeCache.Add(url, body)
	} else {
		body = pokecache
	}
	
	var pokemon registry.Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err!=nil{
		fmt.Println(err)
		return err
	}
	
	if attemptCatch(pokemon.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.CaughtPokemon[pokemon.Name]=pokemon
	} else{
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	
	return nil
}