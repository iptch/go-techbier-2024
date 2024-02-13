package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/iptch/pokedex/pokeapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// a pokemon list item is a pokemon
type PokemonItem pokeapi.PokemonRef

// ### Task 2 ###
// This will fail as long as you do not implement the correct interface
//
// check if type implements interface
var _ list.DefaultItem = (*PokemonItem)(nil)

// ### Task 2 ###
// Through some research on the Go package index we find out that the
// list.DefaultItem interface is composed of the list.Item interface
// and two extra methods Title() and Description(). This is called
// embedding.
// This means we need to define the three methods Title(), Description(),
// and FilterValue(). We do this for a receiver object PokemonItem.

func (i PokemonItem) Title() string {
	// English Titles are tricky, so let's make sure that we get it right
	// by using cases.Title()
	return cases.Title(language.AmericanEnglish).String(i.Name)
}

func (i PokemonItem) Description() string {
	// We do an empty implementation for Description, since we only want
	// to show the Name of the Pokemon in the list view.
	return ""
}

// We enable filtering by the item's name. Go can be pretty concise for
// short methods and functions, as this one-line implementation shows.
func (i PokemonItem) FilterValue() string { return i.Name }
