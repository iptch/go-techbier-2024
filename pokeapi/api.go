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
	// TODO
	//
	// Task 1 is about collecting all PokemonRefs (defined above) from the PokeAPI. Check the
	// `pokemonListUrl` above in your browser to get an idea for the response structure.
	//
	// TASK 1a
	//
	// Create a struct to accomodate the response, similarly to the JSON example in the slides.
	// Using `net/http` from the standard library, do a GET request to the URL above to collect
	// and return the first batch of responses. Use the response `Body` similary to the file handle
	// in the JSON example.
	//
	// Run your code with `go run .`.
	//
	// TASK 1b
	//
	// You should have noticed a field in the JSON response containing the URL of the next chunk
	// of Pokemon. Collect all these responses in a slice (check the slide for examples of loops and
	// slice initialization).
	return nil, fmt.Errorf("not implemented")
}
