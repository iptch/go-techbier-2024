package main

import (
	"fmt"
	"os"

	"github.com/iptch/pokedex/pokeapi"
)

const expectedPokemonCount = 1302

func main() {
	fmt.Println("Downloading pokemon...")
	results, err := pokeapi.GetAllPokemon()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	for _, pokemonRef := range results {
		fmt.Println(pokemonRef.Name)
	}
	fmt.Printf("%d/%d Pokemon listed.\n", len(results), expectedPokemonCount)

	fmt.Println()
	switch len(results) {
	case expectedPokemonCount:
		fmt.Println("Good job! You've completed task 1.")
	case 0:
		fmt.Println("No pokemon listed. Check your code :)")
	default:
		fmt.Println("Looking good! You can move on to task 1b.")
	}
}
