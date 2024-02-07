package pokemon

import (
	"encoding/json"
	"net/http"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

type Pokemon struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	cachedResponse *PokemonResponse
	cachedSprite   string
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

func (p *Pokemon) getResponse() (PokemonResponse, error) {
	if p.cachedResponse != nil {
		return *p.cachedResponse, nil
	}

	resp, err := http.Get(p.Url)
	if err != nil {
		return PokemonResponse{}, err
	}

	p.cachedResponse = new(PokemonResponse)
	err = json.NewDecoder(resp.Body).Decode(p.cachedResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	return *p.cachedResponse, nil
}

func (p *Pokemon) GetTypes() ([]string, error) {
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

func (p *Pokemon) GetAsciiSprite(width int) (string, error) {
	if p.cachedSprite != "" {
		return p.cachedSprite, nil
	}

	response, err := p.getResponse()
	if err != nil {
		return "", err
	}

	spritesUrl := response.Sprites["other"].(map[string]interface{})["official-artwork"].(map[string]interface{})["front_default"].(string)

	flags := aic_package.DefaultFlags()
	flags.Width = width
	flags.Colored = true

	p.cachedSprite, err = aic_package.Convert(spritesUrl, flags)
	return p.cachedSprite, err
}
