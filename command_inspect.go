package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Incorrect number of arguments.")
	}
	name := args[0]

	data, exists := cfg.pokemons[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", data.Name)
	fmt.Println("Height:", data.Height)
	fmt.Println("Weight:", data.Weight)

	/*
	Stats:
	  -hp: 40
	  -attack: 45
	  -defense: 40
	  -special-attack: 35
	  -special-defense: 35
	  -speed: 56
	Types:
	  - normal
	  - flying
	*/
	fmt.Println("Stats:")
	fmt.Println("\t-hp: ", data.)


	return nil
}
