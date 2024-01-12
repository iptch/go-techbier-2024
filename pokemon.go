package main

import (
	"encoding/json"
	"net/http"
)

type AllPokemonResponse struct {
	Results []Pokemon `json:"results"`
	NextUrl string    `json:"next"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonResponse struct {
	Types []struct {
		Slot int         `json:"slot"`
		Type PokemonType `json:"type"`
	} `json:"types"`
}

func GetPokemon() ([]Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon"

	people := []Pokemon{}

	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return []Pokemon{}, err
		}

		var peopleResponse AllPokemonResponse
		err = json.NewDecoder(resp.Body).Decode(&peopleResponse)
		if err != nil {
			return []Pokemon{}, err
		}

		people = append(people, peopleResponse.Results...)

		url = peopleResponse.NextUrl
	}

	return people, nil
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
		return []string{}, nil
	}

    types := []string{}
    for _, type_ := range response.Types {
        types = append(types, type_.Type.Name)
    }

    return types, nil

}
