package pokeapi

import (
	"encoding/json"
	"net/http"
)

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
func GetAllPokemon(n int) ([]PokemonRef, error) {
	results := make([]PokemonRef, 0)

	url := "https://pokeapi.co/api/v2/pokemon"

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
		if len(results) >= n {
			return results[:n], nil
		}

		url = pokemonList.NextUrl
	}

	return results, nil
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
