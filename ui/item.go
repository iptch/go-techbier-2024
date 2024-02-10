package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/iptch/pokedex/pokeapi"
)

// a pokemon list item is a pokemon
type PokemonItem pokeapi.PokeapiRef[pokeapi.Pokemon]

// check if type implements interface
var _ list.DefaultItem = (*PokemonItem)(nil)

// ### Task 2 ###
// Make sure that PokemonItem implements the list.DefaultItem interface
// Check out the documentation for the the interface on pkg.go.dev
//
// If you are stuck, read up on embedding, e.g. here https://go.dev/doc/effective_go#embedding
//
// Think about how you could make the Pokemon list item look good in the list.
// Maybe you want to make sure that the displayed item title is correctly capitalized?
// Hint: check out the golang.org/x/text/cases package
