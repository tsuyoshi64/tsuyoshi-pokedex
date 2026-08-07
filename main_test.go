package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestStartRepl(t *testing.T) {
	// Ensure commands map is initialized
	if commands == nil {
		commands = make(map[string]cliCommand)
	}

	// Register a mock error command specifically for testing error output
	commands["fail"] = cliCommand{
		name:        "fail",
		description: "Triggers a test error",
		callback: func(purls *config) error {
			return errors.New("simulated failure")
		},
	}
	defer delete(commands, "fail")

	cases := []struct {
		name           string
		input          string
		expectedOutput []string
	}{
		{
			name:  "empty input produces prompt only",
			input: "\n",
			expectedOutput: []string{
				"Pokedex >> ",
			},
		},
		{
			name:  "unknown command",
			input: "invalidcmd\n",
			expectedOutput: []string{
				"Unknown command",
			},
		},
		{
			name:  "valid help command execution",
			input: "help\n",
			expectedOutput: []string{
				"Welcome to the Pokedex!",
				"Usage:",
				"help: Displays a help message",
			},
		},
		{
			name:  "multiple commands and whitespace handling",
			input: "   \n  help  \n   foobar \n",
			expectedOutput: []string{
				"Welcome to the Pokedex!",
				"Unknown command",
			},
		},
		{
			name:  "command returning an error",
			input: "fail\n",
			expectedOutput: []string{
				"Error executing command fail: simulated failure",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Redirect os.Stdout to capture fmt.Println output
			oldStdout := os.Stdout
			rPipe, wPipe, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}
			os.Stdout = wPipe

			inBuffer := strings.NewReader(c.input)
			outBuffer := &bytes.Buffer{}

			startRepl(inBuffer, outBuffer)

			// Restore os.Stdout and close pipe writer
			wPipe.Close()
			os.Stdout = oldStdout

			var capturedStdout bytes.Buffer
			_, _ = io.Copy(&capturedStdout, rPipe)
			rPipe.Close()

			// Combine REPL stream buffer and captured stdout
			fullOutput := outBuffer.String() + capturedStdout.String()

			for _, expectedSubstr := range c.expectedOutput {
				if !strings.Contains(fullOutput, expectedSubstr) {
					t.Errorf("expected output to contain %q, but got:\n%s", expectedSubstr, fullOutput)
				}
			}
		})
	}
}

func TestCleanInputEdgeCases(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty input",
			input:    "",
			expected: []string{},
		},
		{
			name:     "spaces, tabs, and newlines",
			input:    " \t  hello \n world  \r\n",
			expected: []string{"hello", "world"},
		},
		{
			name:     "mixed upper and lower case",
			input:    "ExIt HeLp",
			expected: []string{"exit", "help"},
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
