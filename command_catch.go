package main

import (
	"fmt"
	"math/rand"

	"github.com/tsuyoshi64/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pokemon's name is required.")
	}
	if cfg.pokedex == nil {
		cfg.pokedex = make(map[string]pokeapi.Pokemon)
	}

	pokemonName := args[0]

	fmt.Fprintf(cfg.writer, "Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("could not get any data about '%s': %w", pokemonName, err)
	}

	if rand.Intn(300) > pokemon.BaseExperience {
		fmt.Fprintf(cfg.writer, "%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
	} else {
		fmt.Fprintf(cfg.writer, "%s escaped!\n", pokemonName)
	}

	fmt.Fprintln(cfg.writer, "You may now inspect it with the 'inspect' command.")
	return nil
}
