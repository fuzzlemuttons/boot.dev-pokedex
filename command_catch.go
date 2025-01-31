package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Incorrect number of arguments.")
	}
	name := args[0]
	resp, err := cfg.pokeapiClient.PokemonInfo(name)
	if err != nil {
		return err
	}

	fmt.Printf("\nThrowing a Pokeball at %s...\n", name)
	catchChance := rand.Intn(resp.BaseExperience)
	if catchChance < 20 {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokemons[name] = resp
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
