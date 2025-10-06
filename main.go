package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Arundhuti2000/pokedexcli/internal/commands"
) 

func main(){
	
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		if !scanner.Scan() { 
			break 
		}
		text:=scanner.Text()
		inputstring:=commands.CleanInput(text)
		if len(inputstring) > 0 {
			if  value, ok := commands.Mapcommands[inputstring[0]]; ok{
				value.Callback()
			} else{
				fmt.Println("Unknown command")
			}
			// fmt.Println("Your command was:", inputstring[0])
		} else {
			fmt.Println("Your command was: <empty>")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}