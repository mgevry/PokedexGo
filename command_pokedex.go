package main

import (
	"fmt"
)

func commandPokedex(cfg *config, params []string) error {
	fmt.Println("Your Pokedex:")

	for _, info := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", info.Name)
	}

	return nil
}