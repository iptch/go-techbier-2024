package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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

type AppModel struct {
	list list.Model
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// switch msg := msg.(type) {
	// case tea.KeyMsg:
	// 	// Key press switch
	// 	switch msg.String() {
	// 	case "ctrl+c", "q":
	// 		return m, tea.Quit
	// 	case "e":
	// 		// Trigger fight by encountering Pokemon
	// 		return m, tea.Cmd(nil)
	// 	}
	// case tea.WindowSizeMsg:
	// 	//x, y := docStyle().GetFrameSize()
	// 	//m.list.SetSize(msg.Width-x - 64, msg.Height-y)
	// }

	// propagate to the underlying list model
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m AppModel) View() string {
	var sprite string
	var err error

	selectedItem := m.list.SelectedItem()
	if selectedItem != nil {
		sprite, err = selectedItem.(PokemonItem).inner.GetAsciiSprite(60)
		if err != nil {
			panic(err)
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top,
		docStyle().
			Width(80).
			Render(m.list.View()),
		otherStyle().
			Height(m.list.Height()).
			Render(sprite),
	)
}

func main() {
	items := []list.Item{}
	// TODO: display spinner and message as Pokemon are being retrieved
	allPokemon, err := GetAllPokemon()
	if err != nil {
		// TODO(@Selim): We should explain panic in the theory part
		panic(err)
	}

	// create list from pokemon items
	for _, pokemon := range allPokemon {
		items = append(items, PokemonItem{
			inner: pokemon,
		})

	}
	// TODO: need to work on window size, list is too wide
	m := AppModel{list: list.New(items, list.NewDefaultDelegate(), 40, 40)}
	m.list.Title = "Pokédex by ipt"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
