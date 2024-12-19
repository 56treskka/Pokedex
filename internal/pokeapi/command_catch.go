package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/56treskka/pokedexcli/internal/pokecache"
)

func CommandCatch(cfg *Config, cache *pokecache.Cache, args ...string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	pokemon := Pokemon{}
	if err = json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	result := rand.Intn(1000)
	catchThreshold := 1000 - pokemon.BaseExperience

	if result <= catchThreshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.CaughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	fmt.Println()

	return nil
}
