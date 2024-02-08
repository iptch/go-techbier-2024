package pokeapi

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	ditto := PokemonRef{
		Url: "https://pokeapi.co/api/v2/pokemon/ditto",
	}
	types, err := ditto.Get()
	if err != nil {
		t.Fatal("error on fetching types")
	}
	if !reflect.DeepEqual(types.Types[0].Type.Name, "normal") {
		t.Fatal("type mismatch")
	}
}
