package pokemon

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define the model for our TUI. For any type to be a Model, it has to implement
// the Model interface: https://pkg.go.dev/github.com/charmbracelet/bubbletea@v0.25.0#Model
type model struct {
	list              list.Model
	error             error
	downloadCompleted bool
}

// InitialModel instantiates a model with a spinner for the waiting screen,
// a list to hold all retrieved Pokemon items, the initial app and error states.
func InitialModel() model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Pokédex by ipt"

	return model{list: l}
}

// DownloadPokemon will call GetAllPokemon to retrieve Pokémon from the PokéAPI.
// Once the download has completed it sends a downloadCompleted message to the
// bubbles Program.
func (m model) DownloadPokemon(p *tea.Program) {
	c := make(chan []Pokemon)

	go GetAllPokemon(c)

    // create list from Pokémon items
    for downloadedPokemon := range c {
        for _, pokemon := range downloadedPokemon {
            p.Send(newPokemon{pokemon})
        }
    }
    p.Send(downloadCompleted{})
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
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case errMsg:
		m.error = msg
		return m, tea.Quit
	case newPokemon:
		// convert to list items
		var cmds []tea.Cmd
		if len(m.list.Items()) == 0 {
			cmds = append(cmds, m.list.StartSpinner())
		}
		cmds = append(cmds, m.list.InsertItem(len(m.list.Items()), PokemonItem{&msg.pokemon}))
		return m, tea.Batch(cmds...)
	case downloadCompleted:
		m.list.StopSpinner()
		return m, nil
	default:
	}
	// propagate to the underlying list model
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m model) View() string {
	if m.error != nil {
		return fmt.Sprintf("\nThere was an error in the application: %v\n\n", m.error)
	}

	render := m.list.View()

	selectedItem := m.list.SelectedItem()
	if m.list.IsFiltered() && selectedItem != nil {
		sprite, err := selectedItem.(PokemonItem).inner.GetAsciiSprite(60)
		if err != nil {
			// set error and re-render
			m.error = err
			return m.View()
		}
		render = lipgloss.JoinHorizontal(lipgloss.Top, render, sprite)
	}

	return render
}
