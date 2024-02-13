package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/iptch/go-techbier-2024/pokeapi"
)

// a pokemon list item is a pokemon
type PokemonItem pokeapi.PokemonRef

// ### Task 2 ###
// This will fail as long as you do not implement the correct interface
//
// check if type implements interface
var _ list.DefaultItem = (*PokemonItem)(nil)

// ### Task 2 ###
// Make sure that PokemonItem implements the list.DefaultItem interface
// Check out the documentation for the the interface on pkg.go.dev
//
// If you are stuck, read up on embedding, e.g. here:
// https://go.dev/doc/effective_go#embedding
//
// Think about how you could make the Pokemon list item look good in the list.
// Maybe you want to make sure that the displayed item title is correctly capitalized?
// Hint: check out the golang.org/x/text/cases package
