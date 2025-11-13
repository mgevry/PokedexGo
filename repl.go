package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/mgevry/pokedex/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

//takes commands, removes whitespaces before and after command, lowercases all words, puts words of command into slice
func cleanInput(text string) []string {
	fields := strings.Fields(text)
	var words []string
	for _, word := range fields {
		lowerWord := strings.ToLower(word)
		words = append(words, lowerWord)
	}

	return words
}
//runs the REPL
func runREPL(cfg *config) {
	//creating new scanner to read input
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")

		if !scanner.Scan() {
			fmt.Println("Error scanning command")
			continue
		}
		//puts text from input into a variable
		scanText := scanner.Text()
		//cleans the input
		cleanText := cleanInput(scanText)
		//takes the first word of the input
		firstWord := cleanText[0]
		//processes the first word of the input to see if there is a matching command, if there is, store the command name
		command, ok := getCommands()[firstWord]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		//calls the corresponding command
		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Error processing command:", err)
			continue
		}

		
	}
}
//struct containing the name, description and corresponding function for a command on our pokedex cli
type cliCommand struct {
	name string
	description string
	callback func(*config) error
}
//registry of commands for pokedexcli
func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: commandExit,
		},
		
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},

		"map": {
			name: "map",
			description: "Get the next page of locations",
			callback: commandMapf,
		},

		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapb,
		},
	}
}
