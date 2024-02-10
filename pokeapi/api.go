package pokeapi

import (
	"fmt"
)

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

const pokemonListUrl = "https://pokeapi.co/api/v2/pokemon"

// GetAllPokemon fetches all PokemonRefs from the PokeAPI.
func GetAllPokemon() ([]PokemonRef, error) {
	// TODO: Task 1a/1b
	return nil, fmt.Errorf("not implemented")
}
