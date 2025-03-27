package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonListResponse struct {
	Count   int          `json:"count"`
	NextUrl string       `json:"next"`
	Results []PokemonRef `json:"results"`
}

const initialPokemonListUrl = "https://pokeapi.co/api/v2/pokemon"

// GetAllPokemon fetches all PokemonRefs from the PokeAPI.
func GetAllPokemon() ([]PokemonRef, error) {
	currentListUrl := initialPokemonListUrl

	// Slice to collect the results
	pokemonRefs := make([]PokemonRef, 0, 0)

	// Think of this as a while loop :)
	for currentListUrl != "" {

		// TODO: Make an HTTP GET request to the current URL
		response, err := http.Get(currentListUrl)
		if err != nil {
			return nil, err
		}
		defer response.Body.Close()

		// TODO: Parse the response body with JSON into a `PokemonListResponse`
		var pokemonListResponse PokemonListResponse
		if err := json.NewDecoder(response.Body).Decode(&pokemonListResponse); err != nil {
			return nil, err
		}

		// TODO: Update stuff here :)
		currentListUrl = pokemonListResponse.NextUrl
		pokemonRefs = append(pokemonRefs, pokemonListResponse.Results...)

		log.Printf("Collected %d Pokémon...", len(pokemonRefs))
	}

	return pokemonRefs, nil
}

// GetPokemonCount fetches the number of pokemon in the PokeAPI.
func GetPokemonCount() (int, error) {
	response, err := http.Get(initialPokemonListUrl)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var pokemonListResponse PokemonListResponse
	if err := json.NewDecoder(response.Body).Decode(&pokemonListResponse); err != nil {
		return 0, err
	}

	return pokemonListResponse.Count, nil
}
