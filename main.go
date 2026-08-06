package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex >> ")
		if !scanner.Scan() {
			break
		}
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

		cmd, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command %s: %v\n", commandName, err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input: ", err)
		os.Exit(1)
	}
}
