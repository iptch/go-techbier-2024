package main

import (
	"fmt"
	"log"

	"github.com/iptch/go-techbier-2024/pokeapi"
)

func main() {
	count, err := pokeapi.GetPokemonCount()
	if err != nil {
		log.Fatalf("Error getting count: %s\n", err)
	}

	log.Println("Collecting pokemon...")
	results, err := pokeapi.GetAllPokemon()
	if err != nil {
		log.Fatalf("Error collecting Pokemon: %s\n", err)
	}
	for _, pokemonRef := range results {
		fmt.Println(pokemonRef.Name)
	}
	log.Printf("%d/%d Pokemon collected.\n", len(results), count)

	switch len(results) {
	case count:
		log.Println("Good job! You've completed task 1.")
	case 0:
		log.Println("No pokemon were collected.")
	default:
		log.Println("Not quite there yet. Check your code :)")
	}
}
