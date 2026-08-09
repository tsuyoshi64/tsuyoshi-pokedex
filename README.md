# Pokedex CLI

A REPL-based Pokedex CLI written in Go, using the public [PokeAPI](https://pokeapi.co/) to browse locations, find Pokemon, catch them, and inspect your caught collection.

## Installation

### Prerequisites

- Go 1.26.5 or newer
- Internet access, because the CLI fetches live data from PokeAPI

### Clone and Run

```bash
git clone https://github.com/tsuyoshi64/pokedexcli.git
cd pokedexcli
go run .
```

### Build

```bash
go build -o pokedex
./pokedex
```

## Usage

Start the REPL:

```bash
go run .
```

You will see the prompt:

```text
Pokedex >
```

Example session:

```text
Pokedex > help
Welcome to the Pokedex!
Usage:

 - catch:      Catching Pokemon adds them to the user's Pokedex
 - exit:       Exit the Pokedex
 - explore:    Displays a list of all the Pokemon located in the current area
 - help:       Displays a help message
 - inspect:    Prints the name, height, weight, stats and type(s) of a Pokemon
 - map:        Displays the names of the next 20 location areas in the Pokemon world
 - mapb:       Displays the names of the previous 20 location areas in the Pokemon world
 - pokedex:    Prints a list of all the names of the Pokemon the you have caught

Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunnyshore-city-area
sinnoh-pokemon-league-area
oreburgh-mine-1f
oreburgh-mine-b1f
valley-windworks-area
eterna-forest-area
fuego-ironworks-area
mt-coronet-1f-route-207
mt-coronet-2f
mt-coronet-3f
mt-coronet-exterior-snowfall
mt-coronet-exterior-blizzard
mt-coronet-4f
mt-coronet-4f-small-room
mt-coronet-5f
mt-coronet-6f
mt-coronet-1f-from-exterior

Pokedex > explore canalave-city-area
Exploring canalave-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - staryu
 - magikarp
 - gyarados
 - wingull
 - pelipper
 - shellos
 - gastrodon
 - finneon
 - lumineon

Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
You may now inspect it with the 'inspect' command.

Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  -electric

Pokedex > pokedex
Your Pokedex:
  - pikachu

Pokedex > exit
Closing the Pokedex... Goodbye!
```

Catch results are random. If a Pokemon escapes, run `catch <pokemon>` again.

## Commands Reference

| Command | Description |
| --- | --- |
| `help` | Prints the available commands. |
| `exit` | Exits the REPL. |
| `map` | Prints the next 20 location areas from PokeAPI. |
| `mapb` | Prints the previous 20 location areas. If you are already on the first page, it says so. |
| `explore <location>` | Lists Pokemon encounters for a location area, such as `explore canalave-city-area`. |
| `catch <pokemon>` | Attempts to catch a Pokemon by name. Caught Pokemon are added to the in-memory Pokedex. |
| `inspect <pokemon>` | Prints height, weight, stats, and types for a Pokemon you have already caught. |
| `pokedex` | Lists all Pokemon caught during the current REPL session. |

## Architecture Notes

- `main.go` starts the REPL and wires a PokeAPI client with a 5 second HTTP timeout and 1 hour cache interval.
- `repl.go` defines the command registry, input cleanup, and shared REPL config.
- `command_*.go` files implement individual REPL commands.
- `internal/pokeapi` wraps PokeAPI requests and JSON decoding.
- `internal/pokecache` provides a small thread-safe in-memory cache that periodically evicts old responses.
