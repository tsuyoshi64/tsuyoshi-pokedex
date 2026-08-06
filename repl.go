package main

import (
	// "fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	return strings.Fields(lowered)
}
