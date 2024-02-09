package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/iptch/pokedex/pokeapi"
)

// Define the model for our TUI. For any type to be a Model, it has to implement
// the Model interface: https://pkg.go.dev/github.com/charmbracelet/bubbletea@v0.25.0#Model
type model struct {
	list    list.Model
	pokemon *pokeapi.PokemonRef
}

// InitialModel instantiates a model with a spinner for the waiting screen,
// a list to hold all retrieved Pokemon items, the initial app and error states.
func InitialModel() model {
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	l := list.New([]list.Item{}, delegate, 0, 0)
	l.Title = "Pokédex by ipt"

	return model{list: l}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
		return m, m.list.NewStatusMessage(fmt.Sprintf("width=%d, height=%d", msg.Width, msg.Height))
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.list.FilterState() != list.Filtering {
				return m, tea.Quit
			}
		case "enter":
			if m.list.FilterState() != list.Filtering {
				m.pokemon = (*pokeapi.PokemonRef)(m.list.SelectedItem().(*PokemonItem))
				return m, nil
			}
		case "esc":
			if m.pokemon != nil {
				m.pokemon = nil
				// consume escape
				return m, nil
			}
		}
		// leave other keys to fall through to list update
	case NewPokemon:
		var cmds []tea.Cmd
		if len(m.list.Items()) == 0 {
			// download has started
			cmds = append(cmds, m.list.StartSpinner())
		}
		cmds = append(cmds, m.list.InsertItem(len(m.list.Items()), (*PokemonItem)(msg.Pokemon)))
		return m, tea.Batch(cmds...)
	case DownloadCompleted:
		m.list.StopSpinner()
		return m, nil
	default:
		// fall through to list update
	}
	// propagate to the underlying list model
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m model) View() string {
	if m.pokemon != nil {
		return buildViewport(*m.pokemon, m.list.Height(), m.list.Width())
	} else {
		return m.list.View()
	}
}
