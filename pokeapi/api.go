package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonTypeRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonStatRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []PokemonRef `json:"results"`
	NextUrl string       `json:"next"`
}

type PokemonStat struct {
	// no fields used
}

type PokemonType struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name  string `json:"name"`
	Types []struct {
		Slot int            `json:"slot"`
		Type PokemonTypeRef `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int            `json:"base_stat"`
		Stat     PokemonStatRef `json:"stat"`
	} `json:"stats"`
	Sprites map[string]interface{} `json:"sprites"`
}

// GetAllPokemon reads all available Pokémon from the pokeapi incrementally.
// A GET on the url provided returns a list of results and a next URL to perform
// another GET request on for another set of Pokémon.
func GetAllPokemon(n int) ([]PokemonRef, error) {
	results := make([]PokemonRef, 0)

	url := "https://pokeapi.co/api/v2/pokemon"

	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var pokemonList PokemonList
		err = json.NewDecoder(resp.Body).Decode(&pokemonList)
		if err != nil {
			return nil, err
		}

		results = append(results, pokemonList.Results...)
		if len(results) >= n {
			return results[:n], nil
		}

		url = pokemonList.NextUrl
	}

	return results, nil
}

func (p PokemonRef) Get() (*Pokemon, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (p *Pokemon) GetSpriteUrl() (string, error) {
	keys := []string{"other", "official-artwork", "front_default"}

	spritesMap := p.Sprites

	var spritesUrl string
	// ### Task 3 ###
	// You will need to use the index returned by `range`, so replace the placeholder
	for _, key := range keys {

		// ### Task 3 ###
		// You will need to use the value from spritesMap, so replace the placeholder
		_, ok := spritesMap[key]
		if !ok {
			return "", fmt.Errorf("key not found: %s", key)
		}

		// ### Task 3 ###
		// To understand what we are doing here, you might want to check out the response
		// from pokeapi when making a request for a Pokémon: https://pokeapi.co/
		//
		// There are different sprites for every Pokémon, but we want to get the default front.
		// The keys above help us in traversing the response structure to find the right sprite.
		// Basically, the URL we are looking for is nested in:
		// sprites > other > official-artwork > front_default
		// In our traversal of this structure, we need to check that the response actually included
		// what we expect it to, since any of the above parts might be missing.
		//
		// Implement some type assertions for the following:
		//   1. Check if we have reached the end of the keys slice yet
		//   2. If not, check whether the current value is of type map[string]any and update spritesMap
		//  	  with the value for the next iteration, return an error if anything goes wrong
		//   3. If yes, assert that the value is of type string and set our spritesUrl accordingly,
		//      return an error if anything goes wrong

	}

	return spritesUrl, nil
}

func (p *Pokemon) GetAsciiSprite(width int) (string, error) {
	spriteUrl, err := p.GetSpriteUrl()
	if err != nil {
		return "", err
	}

	// ### Task 3 ###
	//
	// We need to convert the Pokemon sprites into ASCII art. We will use the
	// package github.com/TheZoraiz/ascii-image-converter/aic_package.
	//
	// Add the necessary import statements at the top of the file and use the
	// imported package to create an ASCII sprite for our Pokédex.

	// Uncomment these lines once you are ready
	// flags := aic_package.DefaultFlags()
	// flags.Width = width
	// flags.Colored = true

	// Finally, wou will need a second function from the imported package, called Convert().
	// The Convert function takes two parameters: a sprite URL and corresponding flags.
	// Make sure you adjust the return statement correctly.

	return spriteUrl, fmt.Errorf("not implemented yet")

}
