package main

import (
	"fmt"
	"os"

	"github.com/56treskka/pokedexcli/internal/pokeapi"
	"github.com/56treskka/pokedexcli/internal/pokecache"
)

func commandExit(cfg *pokeapi.Config, cache *pokecache.Cache, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
