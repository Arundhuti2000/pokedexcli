package commands

import (
	"fmt"

	"github.com/Arundhuti2000/pokedexcli/internal/registry"
)

func CommandInspect(config *registry.Config, name string) error {
	pokecaught, ok:= config.CaughtPokemon[name]
	if !ok{
		fmt.Println("you have not caught that pokemon")
	} else{
		fmt.Printf("Name:%s\n", pokecaught.Name)
		fmt.Printf("Height:%d\n", pokecaught.Height)
		fmt.Printf("Weight:%d\n", pokecaught.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokecaught.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokecaught.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	}
	return nil
}