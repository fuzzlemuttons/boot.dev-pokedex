package main

import (
	"errors"
	"fmt"
)

func processMapResults(cfg *config, url *string) error {
	resp, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous

	if cfg.prevLocationsURL == nil {
		cfg.nextLocationsURL = nil
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapf(cfg *config, args []string) error {
	return processMapResults(cfg, cfg.nextLocationsURL)
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	return processMapResults(cfg, cfg.prevLocationsURL)
}
