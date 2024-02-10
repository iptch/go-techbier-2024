package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/iptch/pokedex/pokeapi"
	"github.com/iptch/pokedex/ui"
)

// TODO(@Zak): What do you think about keeping the main.go file really lean?
// Also, the code is already getting a bit elaborate at this point. We should sync
// and see how we will present the code to the group such that it is easily
// understandable, would you agree?
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
	c := make(chan []pokeapi.PokeapiRef[pokeapi.Pokemon])

	go pokeapi.GetAllPokemon(c)

	// create list from Pokémon items
	for pokemonRefs := range c {
		for _, pokemonRef := range pokemonRefs {
			pokemonRef := pokemonRef
			p.Send(ui.NewPokemon{Pokemon: pokemonRef})
		}
	}
	p.Send(ui.DownloadCompleted{})
}
