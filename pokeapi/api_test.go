package pokeapi

import (
	"reflect"
	"testing"
)

func TestTypes(t *testing.T) {
	ditto := PokeapiRef[Pokemon]{
		Url: "https://pokeapi.co/api/v2/pokemon/ditto",
	}
	pokemon, err := ditto.Get()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(pokemon.Types[0].Type.Name, "normal") {
		t.Fatal("type mismatch")
	}
}

func TestAscii(t *testing.T) {
	ditto := PokeapiRef[Pokemon]{
		Url: "https://pokeapi.co/api/v2/pokemon/ditto",
	}
	pokemon, err := ditto.Get()
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
