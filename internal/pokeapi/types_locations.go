package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type RespLocation struct {
	ID                int                     `json:"id"`
	Name              string                  `json:"name"`
	PokemonEncounters []RespPokemonEncounters `json:"pokemon_encounters"`
}
type RespPokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type RespPokemonEncounters struct {
	Pokemon RespPokemon `json:"pokemon"`
}

type Pokemon struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	IsDefault      bool        `json:"is_default"`
	Order          int         `json:"order"`
	Weight         int         `json:"weight"`
	Abilities      []Abilities `json:"abilities"`
}
type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Abilities struct {
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
	Ability  Ability `json:"ability"`
}
