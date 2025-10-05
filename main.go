package main

import (
	"bufio"
	"fmt"
	"os"
) 

func main(){
	
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		if !scanner.Scan() { 
			break 
		}
		text:=scanner.Text()
		inputstring:=cleanInput(text)
		if len(inputstring) > 0 {
			if  value, ok := mapcommands[inputstring[0]]; ok{
				value.callback()
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