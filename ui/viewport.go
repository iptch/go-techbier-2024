package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"github.com/iptch/pokedex/pokeapi"
)

func textStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Padding(0, 1).
		Bold(true).
		Italic(true)
}

func typeColors() map[string]lipgloss.Color {
	return map[string]lipgloss.Color{
		"bug":      lipgloss.Color("#D7FF87"),
		"electric": lipgloss.Color("#FDFF90"),
		"fire":     lipgloss.Color("#FF7698"),
		"flying":   lipgloss.Color("#FF87D7"),
		"grass":    lipgloss.Color("#75FBAB"),
		"ground":   lipgloss.Color("#FF875F"),
		"normal":   lipgloss.Color("#929292"),
		"poison":   lipgloss.Color("#7D5AFC"),
		"water":    lipgloss.Color("#00E2C7"),
	}
}

func getTypeColor(typeName string) lipgloss.Color {
	typeColor, ok := typeColors()[typeName]
	if !ok {
		// default
		return lipgloss.Color("#929292")
	}
	return typeColor
}

func buildViewport(pokemon pokeapi.PokemonRef, height, width int) string {
	spriteWidth := 2 * width / 5
	var spriteSection string
	if sprite, err := pokemon.GetAsciiSprite(spriteWidth); err != nil {
		spriteSection = fmt.Sprintf("Could not fetch sprite: %s", err)
	} else {
		spriteSection = sprite
	}

	response, err := pokemon.Get()
	if err != nil {
		return fmt.Sprintf("Error fetching Pokemon: %s", err)
	}

	var descriptionSection string
	descriptionSection += "\n"
	for _, type_ := range response.Types {
		typeName := type_.Type.Name
		descriptionSection += textStyle().
			Foreground(getTypeColor(typeName)).
			Render(strings.ToUpper(typeName))
	}
	descriptionSection += "\n\n"

	progressBar := progress.New(progress.WithDefaultGradient())
	progressBar.PercentFormat = " %.f "

	for _, stat := range response.Stats {
		descriptionSection += progressBar.ViewAs(float64(stat.BaseStat) / 100.0)
		descriptionSection += textStyle().Render(strings.ToUpper(strings.ReplaceAll(stat.Stat.Name, "-", " ")))
		descriptionSection += "\n"
	}

	return lipgloss.JoinHorizontal(lipgloss.Top,
		spriteSection,
		descriptionSection,
	)
}
