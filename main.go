package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type appState int
type downloadComplete struct{ items []list.Item }
type errorMsg struct{ err error }

func (e errorMsg) Error() string { return e.err.Error() }

const (
	downloading appState = iota
	browsing
)

var p *tea.Program

type model struct {
	state   appState
	spinner spinner.Model
	list    list.Model
	error   error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	items := []list.Item{}
	l := list.New(items, list.NewDefaultDelegate(), 40, 40)
	l.Title = "Pokédex by ipt"

	return model{
		state:   downloading,
		spinner: s,
		list:    l,
		error:   nil,
	}

}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) downloadPokemon() {
	allPokemon, err := GetAllPokemon()
	if err != nil {
		p.Send(errorMsg{err})
	}

	// create list from Pokémon items
	items := []list.Item{}
	for _, pokemon := range allPokemon {
		items = append(items, PokemonItem{
			inner: pokemon,
		})
	}

	msgComplete := downloadComplete{
		items: items,
	}
	p.Send(msgComplete)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case downloading:

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "ctrl+c", "q", "esc":
				return m, tea.Quit
			}

		case spinner.TickMsg:
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd

		case errorMsg:
			m.error = msg
			return m, tea.Quit

		case downloadComplete:
			m.list.SetItems(msg.items)
			m.state = browsing
			return m, nil
		}

	case browsing:
		// propagate to the underlying list model
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil

}

func (m model) View() string {
	if m.error != nil {
		return fmt.Sprintf("\nThere was an error in the application: %v\n\n", m.error)
	}

	switch m.state {
	case downloading:
		return "\nWelcome to ipt Pokédex!\n\n\nGetting Pokémon ...\n\n" + m.spinner.View()

	case browsing:
		var sprite string
		var err error

		selectedItem := m.list.SelectedItem()
		if selectedItem != nil {
			sprite, err = selectedItem.(PokemonItem).inner.GetAsciiSprite(60)
			if err != nil {
				m.error = err
				return fmt.Sprintf("\nThere was an error in the application: %v\n\n", m.error)
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
	default:
		return "Unknown state"
	}
}

func main() {
	model := initialModel()
	p = tea.NewProgram(model, tea.WithAltScreen())

	go model.downloadPokemon()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
