package main

import "fmt"

func commandHelp(cfg *config, args []string) error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

`)
	for k, v := range getCommands() {
		fmt.Printf("%v: %v\n", k, v.description)
	}

	return nil
}
