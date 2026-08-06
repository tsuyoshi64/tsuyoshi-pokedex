package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "   spaces   everywhere   ",
			expected: []string{"spaces", "everywhere"},
		},
		{
			input:    "tabs\tseparated\ttext",
			expected: []string{"tabs", "separated", "text"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Content mismatch for input %q: expected %v, got %v", tc.input, tc.expected, actual)
			}
		})
	}
}
