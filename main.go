package main

import (
	"fmt"
	"os"

	"github.com/iptch/pokedex/pokeapi"
)

const expectedPokemonCount = 1302

func main() {
	results, err := pokeapi.GetAllPokemon()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	for _, pokemonRef := range results {
		fmt.Println(pokemonRef.Name)
	}
	fmt.Println()
	if len(results) == expectedPokemonCount {
		fmt.Println("looking good! you've completed task 1.")
	} else {
		fmt.Println("good job, you can move to task 1b.")
	}
}
