package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  hello  WORLD ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  heLlo  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "PIKACHU",
			expected: []string{"pikachu"},
		},
		{
			input:    "Charmander    Bulbasaur       PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expected := c.expected

		if len(actual) != len(expected) {
			t.Errorf("got len: %d, wanted len: %d", len(actual), len(expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("got: %q, wanted: %q", word, expectedWord)
			}
		}
	}
}
