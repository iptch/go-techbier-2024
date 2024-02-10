package pokeapi

import (
	"testing"
)

const dittoUrl = "https://pokeapi.co/api/v2/pokemon/ditto"

func TestAscii(t *testing.T) {
	pokemonRef := PokeapiRef[Pokemon]{Url: dittoUrl}
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

func TestGeneric(t *testing.T) {
	pokemonRef := PokeapiRef[Pokemon]{Url: dittoUrl}
	pokemon, err := pokemonRef.Get()
	if err != nil {
		t.Fatal(err)
	}

	// note the generic type
	typeRef := pokemon.Types[0].Type
	type_, err := typeRef.Get()
	if err != nil {
		t.Fatal(err)
	}

	if type_.Name != "normal" {
		t.Fatal("pokemon type mismatch")
	}
}
