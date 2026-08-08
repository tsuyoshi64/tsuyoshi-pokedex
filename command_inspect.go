package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pokemon's name is required.")
	}

	pokemonName := args[0]

	pokemon, exist := cfg.pokedex[pokemonName]
	if !exist {
		fmt.Fprintln(cfg.writer, "you have not caught that pokemon")
		return nil
	}

	fmt.Fprintf(cfg.writer, "Name: %s\n", pokemon.Name)
	fmt.Fprintf(cfg.writer, "Height: %d\n", pokemon.Height)
	fmt.Fprintf(cfg.writer, "Weight: %d\n", pokemon.Weight)

	fmt.Fprintln(cfg.writer, "Stats:")
	for i := 0; i < len(pokemon.Stats); i++ {
		fmt.Fprintf(cfg.writer, "  -%s: %d\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}

	fmt.Fprintln(cfg.writer, "Types:")
	for i := 0; i < len(pokemon.Types); i++ {
		fmt.Fprintf(cfg.writer, "  -%s\n", pokemon.Types[i].Type.Name)
	}

	return nil
}
