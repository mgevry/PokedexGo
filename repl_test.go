package main

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU   ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "Geralt     of       Rivia",
			expected: []string{"geralt", "of", "rivia"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected: %v does not match Actual: %v", c.expected, actual)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if expectedWord != word {
				t.Errorf("Expected: %s does not match Actual: %s", expectedWord, word)
			}
		}
	}
}