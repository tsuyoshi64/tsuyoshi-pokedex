package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/tsuyoshi64/pokedexcli/internal/pokeapi"
)

func main() {
	startRepl(os.Stdin, os.Stdout)
}

func startRepl(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	cfg := &config{
		writer:        w,
		pokeapiClient: pokeapi.NewClient(5*time.Second, time.Hour),
	}

	for {
		fmt.Fprint(w, "Pokedex >> ")
		if !scanner.Scan() {
			break
		}
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := cleaned[1:]

		cmd, exists := commands[commandName]
		if !exists {
			fmt.Fprintln(w, "Unknown command")
			continue
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Fprintf(w, "Error executing command %s: %v\n", commandName, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(w, "Error reading standard input: %v\n", err)
	}
}
