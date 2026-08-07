package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.next)
	if err != nil {
		return fmt.Errorf("could not get location areas: %w", err)
	}

	cfg.previous = locations.Previous
	cfg.next = locations.Next

	for _, location := range locations.Results {
		fmt.Fprintln(cfg.writer, location.Name)
	}

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.previous == nil {
		fmt.Fprintln(cfg.writer, "you're on the first page")
		return nil
	}

	locations, err := cfg.pokeapiClient.GetLocationAreas(cfg.previous)
	if err != nil {
		return fmt.Errorf("could not get location areas: %w", err)
	}

	cfg.previous = locations.Previous
	cfg.next = locations.Next

	for _, location := range locations.Results {
		fmt.Fprintln(cfg.writer, location.Name)
	}

	return nil
}
