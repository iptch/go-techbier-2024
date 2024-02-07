package pokemon

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define the model for our TUI. For any type to be a Model, it has to implement
// the Model interface: https://pkg.go.dev/github.com/charmbracelet/bubbletea@v0.25.0#Model
type model struct {
	spinner           spinner.Model
	list              list.Model
	error             error
	downloadCompleted bool
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
		spinner: s,
		list:    l,
		error:   nil,
	}
}

// DownloadPokemon will call GetAllPokemon to retrieve Pokémon from the PokéAPI.
// Once the download has completed it sends a downloadCompleted message to the
// bubbles Program.
func (m model) DownloadPokemon(p *tea.Program) {
	c := make(chan []Pokemon)

	go GetAllPokemon(c)

	go func() {
		// create list from Pokémon items
		for downloadedPokemon := range c {
			for _, pokemon := range downloadedPokemon {
				p.Send(newPokemon{pokemon})
			}
		}

		p.Send(downloadCompleted{})
	}()
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
	case newPokemon:
		// convert to list items
		cmd := m.list.InsertItem(len(m.list.Items()), PokemonItem{&msg.pokemon})
		return m, cmd
	case downloadCompleted:
		m.downloadCompleted = true
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
	if len(m.list.Items()) == 0 {
		return "\nWelcome to ipt Pokédex!\n\n\nGetting Pokémon ...   " + m.spinner.View()
	}

	var sprite string
	var err error

	selectedItem := m.list.SelectedItem()
	if m.list.IsFiltered() && selectedItem != nil {
		sprite, err = selectedItem.(PokemonItem).inner.GetAsciiSprite(60)
		if err != nil {
			m.error = err
			return fmt.Sprintf("\nThere was an error in the application: %v\n\n", m.error)
		}
	}

	render := lipgloss.JoinHorizontal(lipgloss.Top,
		docStyle().
			Width(50).
			Render(m.list.View()),
		otherStyle().
			Height(m.list.Height()).
			Render(sprite))

	if !m.downloadCompleted {
		render = lipgloss.JoinVertical(lipgloss.Left, render, m.spinner.View())
	}
	return render
}
