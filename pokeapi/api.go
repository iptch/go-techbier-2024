package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

type PokeapiRef[T any] struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []PokeapiRef[Pokemon] `json:"results"`
	NextUrl string                `json:"next"`
}

type PokemonStat struct {
	// no fields used
}

type PokemonType struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name  string `json:"name"`
	Types []struct {
		Slot int                     `json:"slot"`
		Type PokeapiRef[PokemonType] `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int                     `json:"base_stat"`
		Stat     PokeapiRef[PokemonStat] `json:"stat"`
	} `json:"stats"`
	Sprites map[string]interface{} `json:"sprites"`
}

// GetAllPokemon reads all available Pokémon from the pokeapi incrementally.
// A GET on the url provided returns a list of results and a next URL to perform
// another GET request on for another set of Pokémon.
func GetAllPokemon(c chan []PokeapiRef[Pokemon]) error {
	defer close(c)

	url := "https://pokeapi.co/api/v2/pokemon"

	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var pokemonList PokemonList
		err = json.NewDecoder(resp.Body).Decode(&pokemonList)
		if err != nil {
			return err
		}

		c <- pokemonList.Results

		url = pokemonList.NextUrl
	}

	return nil
}

func (p PokeapiRef[T]) Get() (*T, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon T
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (p *Pokemon) GetSpriteUrl() (string, error) {
	keys := []string{"other", "official-artwork", "front_default"}

	spritesMap := p.Sprites

	var spritesUrl string
	for i, key := range keys {
		value, ok := spritesMap[key]
		if !ok {
			return "", fmt.Errorf("key not found: %s", key)
		}
		if i != len(keys)-1 {
			spritesMap, ok = value.(map[string]any)
			if !ok {
				return "", fmt.Errorf("expected map")
			}
		} else {
			spritesUrl, ok = value.(string)
			if !ok {
				return "", fmt.Errorf("expected string")
			}
		}
	}

	return spritesUrl, nil
}

func (p *Pokemon) GetAsciiSprite(width int) (string, error) {
	spriteUrl, err := p.GetSpriteUrl()
	if err != nil {
		return "", err
	}

	flags := aic_package.DefaultFlags()
	flags.Width = width
	flags.Colored = true

	return aic_package.Convert(spriteUrl, flags)
}
