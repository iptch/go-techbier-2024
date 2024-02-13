package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

// TODO
//
// Bonus bonus task
//
// Don't you feel like there's some repetition in these structs? Go introduced generics some time ago,
// feel free to adapt this code to use generics :)

type PokemonTypeRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonStatRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []PokemonRef `json:"results"`
	NextUrl string       `json:"next"`
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
		Slot int            `json:"slot"`
		Type PokemonTypeRef `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int            `json:"base_stat"`
		Stat     PokemonStatRef `json:"stat"`
	} `json:"stats"`
	Sprites map[string]interface{} `json:"sprites"`
}

// GetAllPokemon reads all available Pokémon from the pokeapi incrementally.
// A GET on the url provided returns a list of results and a next URL to perform
// another GET request on for another set of Pokémon.
func GetAllPokemon(c chan []PokemonRef) error {

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

		// TODO
		//
		// Bonus Task
		//
		// Send results over the channel here. Make sure to close the channel! (remember defer?)

		url = pokemonList.NextUrl
	}

	return nil
}

func (p PokemonRef) Get() (*Pokemon, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
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
