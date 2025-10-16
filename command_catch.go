package main

import (
	"fmt"
	"math/rand"
)


func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("pokemon name not provided")
	}
	pokemon := args[0]
	resp, err := cfg.pokeapiClient.GetPokemonInfo(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	chanceToCatch := resp.BaseExperience
	roll := rand.Intn(1000)
	if roll > chanceToCatch*2 {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokedex[pokemon] = resp
	}else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	return nil
}