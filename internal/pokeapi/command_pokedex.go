package pokeapi

import (
	"fmt"

	"github.com/56treskka/pokedexcli/internal/pokecache"
)

func CommandPokedex(cfg *Config, cache *pokecache.Cache, args ...string) error {
	pokemons := cfg.CaughtPokemon

	for key := range pokemons {
		fmt.Printf(" - %s\n", key)
	}
	fmt.Println()

	return nil
}
