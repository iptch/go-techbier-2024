package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iptch/pokedex/pokemon"
)

// TODO(@Zak): What do you think about keeping the main.go file really lean?
// Also, the code is already getting a bit elaborate at this point. We should sync
// and see how we will present the code to the group such that it is easily
// understandable, would you agree?
func main() {
	model := pokemon.InitialModel()
	program := tea.NewProgram(model, tea.WithAltScreen())

	go model.DownloadPokemon(program)

	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
