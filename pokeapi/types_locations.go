package pokeapi

//struct for location json to unmarshal into, used for map and mapb to list all pokemon locations
type RespShallowLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}

//struct used for single location json to unmarshal into, used for explore to list pokemon in area
type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

