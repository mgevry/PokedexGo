package main

import (
	"strings"
	"fmt"
	"bufio"
)

func cleanInput(text string) []string {
	fields := strings.Fields(text)
	var words []string
	for _, word := range fields {
		lowerWord := strings.ToLower(word)
		words = append(words, lowerWord)
	}

	return words
}

func runREPL(scanner *bufio.Scanner) {
	for {
		fmt.Print("Pokedex >")
		if !scanner.Scan() {
			fmt.Printf("Error scanning command")
			break
		}
		scanText := scanner.Text()
		command := cleanInput(scanText)
		firstWord := command[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}
