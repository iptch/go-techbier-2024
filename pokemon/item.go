package pokemon

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// a pokemon list item is a pokemon
type PokemonItem PokemonResponse

// check if type implements interface
var _ list.Item = (*PokemonItem)(nil)

func (i *PokemonItem) Title() string {
	return cases.Title(language.AmericanEnglish).String(i.Name)
}

func (i *PokemonItem) Description() string {
	types := make([]string, 0, len(i.Types))
	for _, type_ := range i.Types {
		types = append(types, type_.Type.Name)
	}
	return strings.Join(types, ", ")
}

func (i *PokemonItem) FilterValue() string { return i.Name }
