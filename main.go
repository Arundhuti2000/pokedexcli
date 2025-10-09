package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Arundhuti2000/pokedexcli/internal/client"
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
			if  value, ok := client.Mapcommands[inputstring[0]]; ok{
				if inputstring[0] == "explore" && len(inputstring) > 1 {
					value.CallbackWithArg(inputstring[1])
				}else if inputstring[0] == "catch" && len(inputstring) > 1 {
					value.CallbackWithArg(inputstring[1])
				} else if value.Callback != nil {
					value.Callback()
				} else {
					fmt.Println("Command requires an argument")
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
}