package pokeapi

import (
	"fmt"

	"github.com/56treskka/pokedexcli/internal/pokecache"
)

func CommandInspect(cfg *Config, cache *pokecache.Cache, args ...string) error {
	pokemon, ok := cfg.CaughtPokemon[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemon_type := range pokemon.Types {
		fmt.Printf(" - %s\n", pokemon_type.Type.Name)
	}
	fmt.Println()

	return nil
}
