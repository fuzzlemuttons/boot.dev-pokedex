package main

import (
	"fmt"
	"testing"
	"time"

	"boot.dev/sebastian/pokedex/internal/pokeapi"
)

func TestListLocations(t *testing.T) {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	t.Run(fmt.Sprintf("Test case ListLocations"), func(t *testing.T) {
		err := commandMapf(cfg, nil)
		if err != nil {
			t.Error(err)
			return
		}
		cfg.nextLocationsURL = nil
		err = commandMapf(cfg, nil)
		if err != nil {
			t.Error(err)
			return
		}
	})

}
