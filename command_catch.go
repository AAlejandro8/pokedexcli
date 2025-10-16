package main

import (
	"fmt"
	"math/rand"
)


func callbackCatch(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetPokemonInfo(cfg.pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.pokemon)
	chanceToCatch := resp.BaseExperience
	roll := rand.Intn(1000)
	if roll > chanceToCatch*2 {
		fmt.Printf("%s was caught!\n", cfg.pokemon)
		cfg.pokedex[cfg.pokemon] = resp
	}else {
		fmt.Printf("%s escaped!\n", cfg.pokemon)
	}
	return nil
}