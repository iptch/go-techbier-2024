package pokemon

type downloadCompleted struct {
	// empty
}

type newPokemon struct {
	pokemon Pokemon
}

type errMsg error