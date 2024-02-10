package pokeapi

import (
	"encoding/json"
	"net/http"
)

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []PokemonRef `json:"results"`
	NextUrl string       `json:"next"`
}

const pokemonListUrl = "https://pokeapi.co/api/v2/pokemon"

func GetAllPokemon() ([]PokemonRef, error) {
	url := pokemonListUrl

	results := make([]PokemonRef, 0)

	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var pokemonList PokemonList
		err = json.NewDecoder(resp.Body).Decode(&pokemonList)
		if err != nil {
			return nil, err
		}

		results = append(results, pokemonList.Results...)
		url = pokemonList.NextUrl
	}

	return results, nil
}
