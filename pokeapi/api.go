package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tidwall/gjson"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []PokemonRef `json:"results"`
	NextUrl string       `json:"next"`
}

type Pokemon struct {
	Name  string `json:"name"`
	Types []struct {
		Slot int         `json:"slot"`
		Type PokemonType `json:"type"`
	} `json:"types"`
}

// GetAllPokemon reads all available Pokémon from the pokeapi incrementally.
// A GET on the url provided returns a list of results and a next URL to perform
// another GET request on for another set of Pokémon.
func GetAllPokemon(c chan []PokemonRef) error {
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

func (p *PokemonRef) Get() (Pokemon, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (p *PokemonRef) GetAsciiSprite(width int) (string, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	json, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if !gjson.ValidBytes(json) {
		return "", fmt.Errorf("invalid json body")
	}

	spritesUrl, ok := gjson.GetBytes(json, "sprites.other.official-artwork.front_default").Value().(string)
	if !ok {
		return "", fmt.Errorf("invalid json value for sprite")
	}
	// spritesUrl := response.Sprites["other"].(map[string]interface{})["official-artwork"].(map[string]interface{})["front_default"].(string)

	flags := aic_package.DefaultFlags()
	flags.Width = width
	flags.Colored = true

	return aic_package.Convert(spritesUrl, flags)
}
