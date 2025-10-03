package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
) 

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string{
	textLower:= strings.ToLower(text)
	words:= strings.Fields(textLower)
	return words
}

func commandExit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return  nil
}

func main(){
	mapexit:=map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		if !scanner.Scan() { 
			break 
		}
		text:=scanner.Text()
		inputstring:=cleanInput(text)
		if len(inputstring) > 0 {
			fmt.Println("Your command was:", inputstring[0])
		} else {
			fmt.Println("Your command was: <empty>")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}