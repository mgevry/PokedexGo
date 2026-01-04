package main

import (
	"time"

	"github.com/mgevry/pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.PokemonInfo),
	}

	runREPL(cfg)
}
