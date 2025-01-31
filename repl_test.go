package main

import (
	"testing"
	"time"

	"boot.dev/sebastian/pokedex/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg, "map")
}
