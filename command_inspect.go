package main

import (
	"fmt"
)

func commandInspect(cfg *config, params []string) error {
	pokemonName := params[0]
	pokeInfo, exists := cfg.caughtPokemon[pokemonName]
	
	if exists == false {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name:", pokeInfo.Name)
		fmt.Println("Height:", pokeInfo.Height)
		fmt.Println("Weight:", pokeInfo.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokeInfo.Stats {
			fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		fmt.Println("Types:")
		for _, t := range pokeInfo.Types {
			fmt.Printf("	- %s\n", t.Type.Name)
		}
	}	
	return nil
}