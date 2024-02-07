package pokemon

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PokemonItem struct {
	inner *Pokemon
}

func (i PokemonItem) Title() string {
	return cases.Title(language.AmericanEnglish).String(i.inner.Name)
}

func (i PokemonItem) Description() string {
	return ""
}

func (i PokemonItem) FilterValue() string { return i.inner.Name }
