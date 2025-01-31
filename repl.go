package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"boot.dev/sebastian/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokemons         map[string]pokeapi.Pokemon
}

func startRepl(cfg *config, overrideCommand string) {
	if len(overrideCommand) > 0 {
		cmd, ok := getCommands()[overrideCommand]
		if !ok {
			fmt.Println("Unknown command")
			return
		}
		err := cmd.callback(cfg, make([]string, 0))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		if len(words) > 0 {
			cmdStr := words[0]
			cmd, ok := getCommands()[cmdStr]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			err := cmd.callback(cfg, words[1:])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	str := strings.ToLower(text)
	words := strings.Fields(str)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
	args        []string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokedex",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
}
