package main

import (
	"math/rand"
	"fmt"
)

func commandCatch(cfg *config, params []string) error {
	pokemonName := params[0]
	msg := fmt.Sprintf("Throwing a Pokeball at %s...", pokemonName)
	fmt.Println(msg)

	pokemonStats, err := cfg.pokeapiClient.PokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	baseExperience := pokemonStats.BaseExperience
	catchChance := (1000 - baseExperience)/10

	catchNumber := rand.Intn(101)

	if catchNumber < catchChance {
		fmt.Println(pokemonName, "was caught!")
		cfg.caughtPokemon[pokemonName] = pokemonStats

	}else{
		fmt.Println(pokemonName, "escaped!")
	}
	
	return nil

}