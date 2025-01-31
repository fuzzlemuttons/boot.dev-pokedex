package main

import (
	"fmt"
	"testing"
	"time"

	"boot.dev/sebastian/pokedex/internal/pokeapi"
)

func TestCatch(t *testing.T) {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokemons:      map[string]pokeapi.Pokemon{},
	}

	t.Run(fmt.Sprintf("Test case ListLocations"), func(t *testing.T) {
		for i := 0; i < 10; i++ {
			err := commandCatch(cfg, []string{"pikachu"})
			if err != nil {
				t.Error(err)
				return
			}
		}
	})

}
