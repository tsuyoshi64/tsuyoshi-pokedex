package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScannerLoopBehavior(t *testing.T) {
	simulatedInput := strings.Join([]string{
		"hello world",                        // Standard input
		"   Charmander Bulbasaur PIKACHU   ", // Mixed case with trailing/leading spaces
		"",                                   // Blank line (should be skipped)
		"tabs\tseparated\twords",             // Tab characters as whitespace
		"   ",                                // Only whitespace (should be skipped)
		"123456 catch_them_all!",             // Numbers and symbols
		"Squirtle",                           // Single word command
		"pokedex   ",                         // Word with trailing spaces
	}, "\n") + "\n"

	reader := strings.NewReader(simulatedInput)
	scanner := bufio.NewScanner(reader)

	expectedFirstWords := []string{
		"hello",
		"charmander",
		"tabs",
		"123456",
		"squirtle",
		"pokedex",
	}

	outputCount := 0

	for scanner.Scan() {
		input := scanner.Text()
		cleanedWords := cleanInput(input)

		if len(cleanedWords) == 0 {
			continue
		}

		if outputCount >= len(expectedFirstWords) {
			t.Fatalf("Received more valid commands than expected. Extra input parsed: %q", input)
		}

		expected := expectedFirstWords[outputCount]
		actual := cleanedWords[0]

		if actual != expected {
			t.Errorf("Line match fail at index %d: expected first word %q, got %q (Original raw input: %q)",
				outputCount, expected, actual, input)
		}

		outputCount++
	}

	if err := scanner.Err(); err != nil {
		t.Errorf("Scanner encountered an unexpected runtime error: %v", err)
	}

	if outputCount != len(expectedFirstWords) {
		t.Errorf("Count mismatch: Expected to process %d commands, but only processed %d",
			len(expectedFirstWords), outputCount)
	}
}
