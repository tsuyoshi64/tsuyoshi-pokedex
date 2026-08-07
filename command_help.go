package main

import (
	"fmt"
	"sort"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		cmd := commands[k]
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
