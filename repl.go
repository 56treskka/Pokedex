package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/56treskka/pokedexcli/internal/pokeapi"
	"github.com/56treskka/pokedexcli/internal/pokecache"
)

func startRepl() error {

	scanner := bufio.NewScanner(os.Stdin)
	cfg := pokeapi.Config{
		Next:          "https://pokeapi.co/api/v2/location-area/",
		Previous:      "",
		CaughtPokemon: map[string]pokeapi.Pokemon{},
	}

	cache := pokecache.NewCache(5 * time.Second)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			words := cleanInput(scanner.Text())
			if len(words) == 0 {
				continue
			}

			commandName := words[0]
			args := []string{}
			if len(words) > 1 {
				args = words[1:]
			}

			commands := getCommands()
			if command, ok := commands[commandName]; ok {
				err := command.callback(&cfg, cache, args...)
				if err != nil {
					return err
				}
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

func cleanInput(text string) []string {
	slice := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return slice
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokeapi.Config, cache *pokecache.Cache, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    pokeapi.CommandMap,
		},

		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    pokeapi.CommandMapb,
		},

		"explore": {
			name:        "explore <area_name>",
			description: "Displays a list of all the Pokémon located in the area",
			callback:    pokeapi.CommandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Tries to catch Pokemon and displays the result",
			callback:    pokeapi.CommandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Displays the name, height, weight, stats and type(s) of the Pokemon",
			callback:    pokeapi.CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays a list of all the names of the Pokemon the user has caught",
			callback:    pokeapi.CommandPokedex,
		},
	}
}
