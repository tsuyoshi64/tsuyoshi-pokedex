package main

import (
	"fmt"
	"sort"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Fprintln(cfg.writer, "Welcome to the Pokedex!")
	fmt.Fprintln(cfg.writer, "Usage:")
	fmt.Fprintln(cfg.writer)

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		cmd := commands[k]
		fmt.Fprintf(cfg.writer, " - %s:\t%s\n", cmd.name, cmd.description)
	}
	return nil
}
