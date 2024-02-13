package pokeapi

import (
	"testing"
)

const dittoUrl = "https://pokeapi.co/api/v2/pokemon/ditto"

func TestSpriteUrl(t *testing.T) {
	pokemonRef := PokemonRef{Url: dittoUrl}
	pokemon, err := pokemonRef.Get()
	if err != nil {
		t.Fatal(err)
	}
	spriteUrl, err := pokemon.GetSpriteUrl()
	if err != nil {
		t.Fatal(err)
	}
	if spriteUrl != "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/132.png" {
		t.Fatal("url mismatch")
	}
}
