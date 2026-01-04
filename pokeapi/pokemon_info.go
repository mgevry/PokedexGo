package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) PokemonInfo(pokemonName string) (PokemonInfo, error){
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		cachePokemonInfo := PokemonInfo{}
		err := json.Unmarshal(val, &cachePokemonInfo)
		if err != nil {
			return PokemonInfo{}, err
		}
		return cachePokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokemonInfo := PokemonInfo{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return PokemonInfo{}, err
	}
	return pokemonInfo, nil
}