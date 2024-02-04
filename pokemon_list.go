package main

import (
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func docStyle() lipgloss.Style {
	return lipgloss.NewStyle().Margin(1, 2)
}

func otherStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
}

type PokemonItem struct {
	inner Pokemon
}

func (i PokemonItem) Title() string {
	return cases.Title(language.AmericanEnglish).String(i.inner.Name)
}

func (i PokemonItem) Description() string {
	return ""
}

func (i PokemonItem) FilterValue() string { return i.inner.Name }
