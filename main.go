package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func docStyle() lipgloss.Style {
	return lipgloss.NewStyle().Margin(1, 2)
}

func otherStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("228")).
		BorderBackground(lipgloss.Color("63")).
		BorderTop(true).
		BorderLeft(true)
}

type PokemonItem struct {
	inner Pokemon
}

func (i PokemonItem) Title() string { return strings.Title(i.inner.Name) }
func (i PokemonItem) Description() string {
	return ""
}
func (i PokemonItem) FilterValue() string { return i.inner.Name }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		x, y := docStyle().GetFrameSize()
		m.list.SetSize(msg.Width-x-32, msg.Height-y)
	}

    // propagate to the underlying list model
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	speciesNames, err := m.list.SelectedItem().(PokemonItem).inner.GetTypes()
	if err != nil {
		panic(err)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, docStyle().Render(m.list.View()), otherStyle().Height(m.list.Height()).Width(32).Render(strings.Join(speciesNames, ", ")))
}

func main() {
	items := []list.Item{}
	people, err := GetPokemon()
	if err != nil {
		panic(err)
	}
	for _, person := range people {
		items = append(items, PokemonItem{
			inner: person,
		})

	}
	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Star Wars People"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

