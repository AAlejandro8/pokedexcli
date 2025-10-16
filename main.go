package main

import "github.com/AAlejandro8/pokedexcli/internal/pokeapi"



func main() {
	cfg := &config {
		pokeapiClient: pokeapi.NewClient(),
		pokedex: make(Pokedex),
	}
	startRepl(cfg)
}