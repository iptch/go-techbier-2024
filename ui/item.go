package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/iptch/pokedex/pokeapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// a pokemon list item is a pokemon
type PokemonItem pokeapi.PokeapiRef[pokeapi.Pokemon]

// check if type implements interface
var _ list.Item = (*PokemonItem)(nil)

func (i PokemonItem) Title() string {
	return cases.Title(language.AmericanEnglish).String(i.Name)
}

func (i PokemonItem) Description() string {
	return ""
}

func (i PokemonItem) FilterValue() string { return i.Name }
