package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Incorrect number of arguments.")
	}
	location := args[0]
	resp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("\nExploring %s...\n", location)
	for _, loc := range resp.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}

	return nil
}
