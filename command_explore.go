package main

import (
	"fmt"
)

func commandExplore(cfg *config, params []string) error {
	location := params[0]
	locationArea, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}
	for _, pe := range locationArea.PokemonEncounters {
		fmt.Println(pe.Pokemon.Name)
	}
	return nil
} 