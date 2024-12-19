package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/56treskka/pokedexcli/internal/pokecache"
)

type Config struct {
	Next          string
	Previous      string
	CaughtPokemon map[string]Pokemon
}

type Page struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CommandMap(cfg *Config, cache *pokecache.Cache, args ...string) error {
	if cfg.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	body, ok := cache.Get(cfg.Next)

	if !ok {
		resp, err := http.Get(cfg.Next)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(cfg.Next, body)
	}

	area := Page{}
	if err := json.Unmarshal(body, &area); err != nil {
		return err
	}

	cfg.Next = area.Next
	cfg.Previous = area.Previous

	for _, result := range area.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()
	return nil
}

func CommandMapb(cfg *Config, cache *pokecache.Cache, args ...string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	body, ok := cache.Get(cfg.Previous)
	if !ok {
		resp, err := http.Get(cfg.Previous)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cache.Add(cfg.Previous, body)
	}

	area := Page{}
	if err := json.Unmarshal(body, &area); err != nil {
		return err
	}

	cfg.Next = area.Next
	cfg.Previous = area.Previous

	for _, result := range area.Results {
		fmt.Println(result.Name)
	}
	fmt.Println()
	return nil
}
