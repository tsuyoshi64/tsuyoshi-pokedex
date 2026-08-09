package main

import (
	"io"
	"strings"

	"github.com/tsuyoshi64/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokémon located in the current area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catching Pokemon adds them to the user's Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Prints the name, height, weight, stats and type(s) of a Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of all the names of the Pokemon the you have caught",
			callback:    commandPokedex,
		},
	}
}

type config struct {
	pokeapiClient pokeapi.Client
	previous      *string
	next          *string
	writer        io.Writer
	pokedex       map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}
