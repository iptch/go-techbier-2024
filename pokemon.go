package main

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

func GetAllPokemon() ([]Pokemon, error) {
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

		pokemon = append(pokemon, allPokemonResponse.Results...)

		url = allPokemonResponse.NextUrl
	}

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
    //resp, err := http.Get(spritesUrl)
    //if err != nil {
    //    return "", err
    //}

    //img, _, err := image.Decode(resp.Body)
    //if err != nil {
    //    return "", err
    //}

    //convertOptions := convert.DefaultOptions
	//convertOptions.FixedWidth = width

    //converter := convert.NewImageConverter()
    //return converter.Image2ASCIIString(img, &convertOptions), nil

    flags := aic_package.DefaultFlags()
    flags.Width = width
    flags.Colored = true

    return aic_package.Convert(spritesUrl, flags)
}
