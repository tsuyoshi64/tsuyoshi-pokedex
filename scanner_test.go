package main

import (
	"bufio"
	"sort"
	"strings"
	"testing"
)

func TestScannerAndRegistryValidation(t *testing.T) {
	simulatedInput := strings.Join([]string{
		"HELP",            // Upper case conversion test
		"   exiT   ",      // Leading/trailing spaces + mixed case
		"",                // Empty string line
		"invalid-command", // Unregistered token
		"   ",             // Pure space string line
		"exit help",       // Multi-argument entry (should only match "exit")
	}, "\n") + "\n"

	reader := strings.NewReader(simulatedInput)
	scanner := bufio.NewScanner(reader)

	type expectedMatch struct {
		rawToken    string
		shouldExist bool
		expectedKey string
	}

	expectedFlow := []expectedMatch{
		{rawToken: "HELP", shouldExist: true, expectedKey: "help"},
		{rawToken: "   exiT   ", shouldExist: true, expectedKey: "exit"},
		{rawToken: "invalid-command", shouldExist: false, expectedKey: "invalid-command"},
		{rawToken: "exit help", shouldExist: true, expectedKey: "exit"},
	}

	flowIndex := 0

	for scanner.Scan() {
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		if flowIndex >= len(expectedFlow) {
			t.Fatalf("Processed more non-empty commands than expected array bounds. Extra line: %q", input)
		}

		target := expectedFlow[flowIndex]
		commandName := cleaned[0] // The REPL evaluates the first token slice element

		if commandName != target.expectedKey {
			t.Errorf("Parsing/cleaning logic failure: expected key token %q, processed %q from raw line %q",
				target.expectedKey, commandName, input)
		}

		_, exists := commands[commandName]
		if exists != target.shouldExist {
			t.Errorf("Registry state failure for command key %q: expected inclusion to be %v, got %v",
				commandName, target.shouldExist, exists)
		}

		flowIndex++
	}

	if err := scanner.Err(); err != nil {
		t.Errorf("Unexpected scanner stream truncation error: %v", err)
	}

	if flowIndex != len(expectedFlow) {
		t.Errorf("Test failed to check all sequential stream assertions. Evaluated %d out of %d entries",
			flowIndex, len(expectedFlow))
	}
}

func TestRegistrySortingLogic(t *testing.T) {
	var keys []string
	for k := range commands {
		keys = append(keys, k)
	}

	if len(keys) < 2 {
		t.Fatalf("Global command registry map contains insufficient entries for sorting validation: %v", keys)
	}

	sort.Strings(keys)

	for i := 0; i < len(keys)-1; i++ {
		if keys[i] > keys[i+1] {
			t.Errorf("Sorting calculation failure: key %q incorrectly sequenced before %q", keys[i], keys[i+1])
		}
	}

	if keys[0] != "exit" || keys[1] != "help" {
		t.Errorf("Dynamic alphabet sorting returned bad current index arrays: expected [exit, help], got %v", keys)
	}
}
