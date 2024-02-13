package ui

import "github.com/iptch/go-techbier-2024/pokeapi"

type DownloadCompleted struct {
	// empty
}

type NewPokemon struct {
	Pokemon pokeapi.PokeapiRef[pokeapi.Pokemon]
}
