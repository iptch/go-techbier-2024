package ui

import "github.com/iptch/pokedex/pokeapi"

type DownloadCompleted struct {
	// empty
}

type NewPokemon struct {
	Pokemon *pokeapi.PokeapiRef[pokeapi.Pokemon]
}
