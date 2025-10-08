package main

import (
	"testing"

	"github.com/Arundhuti2000/pokedexcli/internal/commands"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
	actual := commands.CleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	if len(actual) != len(c.expected) {
            t.Errorf("Expected length %d, got %d", len(c.expected), len(actual))
            continue
        }
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		 if word != expectedWord {
        	t.Errorf("Expected: %s, Actual: %s", expectedWord, word)
    	} else {
			t.Logf("PASS: input=%q output=%v", c.input, actual)
		}
		}
	}
}