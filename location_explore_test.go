package main

import (
	"fmt"
	"testing"
	"time"

	"boot.dev/sebastian/pokedex/internal/pokeapi"
)

func TestExplore(t *testing.T) {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	t.Run(fmt.Sprintf("Test case ListLocations"), func(t *testing.T) {
		err := commandExplore(cfg, []string{"mt-coronet-6f"})
		if err != nil {
			t.Error(err)
			return
		}

	})

}
