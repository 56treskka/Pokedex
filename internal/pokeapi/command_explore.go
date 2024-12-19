package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/56treskka/pokedexcli/internal/pokecache"
)

type Area struct {
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func CommandExplore(cfg *Config, cache *pokecache.Cache, args ...string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Exploring %s...\n", args[0])

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	area := Area{}

	if err = json.Unmarshal(body, &area); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range area.Encounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	fmt.Println()

	return nil
}
