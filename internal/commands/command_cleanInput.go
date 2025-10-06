package commands

import "strings"

func CleanInput(text string) []string{
	textLower:= strings.ToLower(text)
	words:= strings.Fields(textLower)
	return words
}