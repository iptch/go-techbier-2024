package pokemon

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// define some custom types needed for our TUI
type appState int
type downloadCompleted struct{ items []list.Item }

// we will define a custom error message type for error handling ...
type errMsg struct{ err error }

// ... and implement the Error interface for it
func (e errMsg) Error() string { return e.err.Error() }

const (
	downloading appState = iota // iota is 0 here
	browsing                    // implicitly uses iota, which is now 1
)

// Define the model for our TUI. For any type to be a Model, it has to implement
// the Model interface: https://pkg.go.dev/github.com/charmbracelet/bubbletea@v0.25.0#Model
type model struct {
	state   appState
	spinner spinner.Model
	list    list.Model
	error   error
}

// InitialModel instantiates a model with a spinner for the waiting screen,
// a list to hold all retrieved Pokemon items, the initial app and error states.
func InitialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("76"))

	items := []list.Item{}
	l := list.New(items, list.NewDefaultDelegate(), 50, 50)
	l.Title = "Pokédex by ipt"

	return model{
		state:   downloading,
		spinner: s,
		list:    l,
		error:   nil,
	}
}

// DownloadPokemon will call GetAllPokemon to retrieve Pokémon from the PokéAPI.
// Once the download has completed it sends a downloadCompleted message to the
// bubbles Program.
func (m model) DownloadPokemon(p *tea.Program) {
	allPokemon, err := GetAllPokemon()
	if err != nil {
		p.Send(errMsg{err})
	}

	// create list from Pokémon items
	items := []list.Item{}
	for _, pokemon := range allPokemon {
		items = append(items, PokemonItem{
			inner: pokemon,
		})
	}

	msgComplete := downloadCompleted{
		items: items,
	}
	p.Send(msgComplete)
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
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

		case errMsg:
			m.error = msg
			return m, tea.Quit

		case downloadCompleted:
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

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m model) View() string {
	if m.error != nil {
		return fmt.Sprintf("\nThere was an error in the application: %v\n\n", m.error)
	}

	switch m.state {
	case downloading:
		return "\nWelcome to ipt Pokédex!\n\n\nGetting Pokémon ...   " + m.spinner.View()

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
				Width(50).
				Render(m.list.View()),
			otherStyle().
				Height(m.list.Height()).
				Render(sprite),
		)
	default:
		return "Unknown state"
	}
}
