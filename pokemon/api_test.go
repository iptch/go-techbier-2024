package pokemon

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	ditto := Pokemon{
		Url: "https://pokeapi.co/api/v2/pokemon/ditto",
	}
	types, err := ditto.GetTypes()
	if err != nil {
		t.Fatal("error on fetching types")
	}
	if !reflect.DeepEqual(types, []string{"normal"}) {
		t.Fatal("type mismatch")
	}
}
