package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// ### Task 3 ###
	//
	// We need to convert the Pokemon sprites into ASCII art. We will use the
	// package github.com/TheZoraiz/ascii-image-converter/aic_package.
	//
	// Add the necessary import statements at the top of the file and use the
	// imported package to create an ASCII sprite for our Pokédex.
	//
	// Hint: you will need to use a function from the aic_package called DefaultFlags().
	// On the return type from that function call, you will set a couple of fields, called
	// Width and Colored. Where do we get the width from? Do we want the sprite to be
	// colored?
	// Finally, wou will need a second function from the imported package, called Convert().
	// Make sure you adjust the return statement correctly.

	return spriteUrl, fmt.Errorf("not implemented yet")

}
