package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Fprintln(cfg.writer, "You have not caught any Pokemon yet")
		return nil
	}

	fmt.Fprintln(cfg.writer, "Your Pokedex:")

	for pokemonName, _ := range cfg.pokedex {
		fmt.Fprintf(cfg.writer, "  - %s\n", pokemonName)
	}
	return nil
}
