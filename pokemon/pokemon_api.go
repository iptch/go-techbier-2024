package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type AllPokemonResponse struct {
	Results []Pokemon `json:"results"`
	NextUrl string    `json:"next"`
}

type PokemonResponse struct {
	Types []struct {
		Slot int         `json:"slot"`
		Type PokemonType `json:"type"`
	} `json:"types"`
	Sprites map[string]interface{}
}

// GetAllPokemon reads all available Pokémon from the pokeapi incrementally.
// A GET on the url provided returns a list of results and a next URL to perform
// another GET request on for another set of Pokémon.
func GetAllPokemon(c chan []Pokemon) ([]Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon"

	pokemon := []Pokemon{}

	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return []Pokemon{}, err
		}

		var allPokemonResponse AllPokemonResponse
		err = json.NewDecoder(resp.Body).Decode(&allPokemonResponse)
		if err != nil {
			return []Pokemon{}, err
		}

		c <- allPokemonResponse.Results

		url = allPokemonResponse.NextUrl
	}

	close(c)

	return pokemon, nil
}

func (p Pokemon) getResponse() (PokemonResponse, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return PokemonResponse{}, err
	}

	var pokemonResponse PokemonResponse
	err = json.NewDecoder(resp.Body).Decode(&pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	return pokemonResponse, nil
}

func (p Pokemon) GetTypes() ([]string, error) {
	response, err := p.getResponse()
	if err != nil {
		return []string{}, err
	}

	types := []string{}
	for _, type_ := range response.Types {
		types = append(types, type_.Type.Name)
	}

	return types, nil

}

func (p Pokemon) GetAsciiSprite(width int) (string, error) {
	response, err := p.getResponse()
	if err != nil {
		return "", err
	}

	spritesUrl := response.Sprites["other"].(map[string]interface{})["official-artwork"].(map[string]interface{})["front_default"].(string)

	flags := aic_package.DefaultFlags()
	flags.Width = width
	flags.Colored = true

	return aic_package.Convert(spritesUrl, flags)
}
