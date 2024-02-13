package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/iptch/go-techbier-2024/pokeapi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// a pokemon list item is a pokemon
type PokemonItem pokeapi.PokemonRef

// check if type implements interface
var _ list.Item = (*PokemonItem)(nil)

func (i PokemonItem) Title() string {
	return cases.Title(language.AmericanEnglish).String(i.Name)
}

func (i PokemonItem) Description() string {
	return ""
}

func (i PokemonItem) FilterValue() string { return i.Name }
