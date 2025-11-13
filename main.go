package main

import (
	"time"

	"github.com/mgevry/pokedex/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	runREPL(cfg)
}
