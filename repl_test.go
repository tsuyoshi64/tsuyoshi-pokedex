package main

import (
	"slices"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only whitespace",
			input:    "   ",
			expected: []string{},
		},
		{
			name:     "tabs and newlines",
			input:    "hello\tworld\nfrom\r\ngo",
			expected: []string{"hello", "world", "from", "go"},
		},
		{
			name:     "single word with leading/trailing spaces",
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			name:     "multiple spaces between words",
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed capitalization",
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "numbers and punctuation",
			input:    "Pikachu #25!  level 100",
			expected: []string{"pikachu", "#25!", "level", "100"},
		},
		{
			name:     "unicode / non-ASCII characters",
			input:    "  Café   Latté  ",
			expected: []string{"café", "latté"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if !slices.Equal(actual, c.expected) {
				t.Errorf("cleanInput(%q) = %v, expected %v", c.input, actual, c.expected)
			}
		})
	}
}

