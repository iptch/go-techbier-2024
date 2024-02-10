package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	// ### Task 3 ###
	// We import the aic_package from Github. It is a convention
	// to leave an empty line between imports from the standard
	// library and those not belonging to the standard library.
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

	// ### Task 3 ###
	// With the package imported, we can investigate what we would use
	// DefaultFlags() for. We add a call to the function by prefixing
	// the function name with the package name.
	flags := aic_package.DefaultFlags()

	// ### Task 3 ###
	// The return type is a Flags struct. Our task was to set two fields
	// on the struct: Width and Colored. We get the width through the
	// function parameters. Whether you want the sprite to be colored or
	// not is personal preference, so you may choose whatever you prefer.
	// We like colors.
	flags.Width = width
	flags.Colored = true

	// ### Task 3 ###
	// Instead of returning an error stating that this function is not
	// implemented, we can now convert the ASCII art sprite. The hint
	// mentioned using a function called Convert(). We do not have to
	// assign a variable to the function but can return its output right
	// away. It takes the URL to our sprite and the previously set flags
	// and returns our ASCII art sprite.
	return aic_package.Convert(spriteUrl, flags)
}
