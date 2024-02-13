package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iptch/go-techbier-2024/pokeapi"
	"github.com/iptch/go-techbier-2024/ui"
)

func main() {
	numPokemon := flag.Int("n", 100, "number of pokemon to download")
	flag.Parse()

	fmt.Println("downloading pokemon...")

	pokemon, err := pokeapi.GetAllPokemon(*numPokemon)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	model := ui.InitialModel(pokemon)
	program := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
