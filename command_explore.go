package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("A location area name is required.")
	}
	areaPokemons, err := cfg.pokeapiClient.GetLocationAreaPokemons(args[0])
	if err != nil {
		return fmt.Errorf("could not get any pokemon data: %w", err)
	}

	fmt.Fprintf(cfg.writer, "Exploring %s...\n", areaPokemons.Name)
	fmt.Fprintln(cfg.writer, "Found Pokemon:")

	for _, pokemon := range areaPokemons.PokemonEncounters {
		fmt.Fprintf(cfg.writer, " - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
