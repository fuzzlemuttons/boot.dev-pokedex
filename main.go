package main

import (
	"boot.dev/sebastian/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 1*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemons:      make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg, "")
}
