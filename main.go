package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iptch/pokedex/ui"
)

func main() {
	model := ui.InitialModel()
	program := tea.NewProgram(model, tea.WithAltScreen())

	go DownloadPokemon(program)

	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}

// DownloadPokemon will call GetAllPokemon to retrieve Pokémon from the PokéAPI.
// Once the download has completed it sends a downloadCompleted message to the
// bubbles Program.
func DownloadPokemon(p *tea.Program) {
	// TODO
	//
	// Bonus Task
	//
	// Here we want to download all the Pokemon in a goroutine. To do this, create
	// a channel here with `make`, and pass it to the GetAllPokemon function from the PokeAPI. You'll need
	// to adapt GetAllPokemon to send messages over the channel, and close the channel when completed.
	//
	// Then in this function go through the elements of this channel, and for each one, send a `NewPokemon` message
	// to the bubbletea program.

	// Notify the program that download has completed.
	p.Send(ui.DownloadCompleted{})
}
