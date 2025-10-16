package main

import (
	"errors"
	"fmt"
)

func callBackPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) < 1 {
		return errors.New("no pokemon in the pokedex")
	}
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}