package main

import "fmt"

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Pokemon name not provided")
	}
	pokemonToInspect := args[0]
	pokemon, ok := cfg.pokedex[pokemonToInspect]
	if !ok {
		return fmt.Errorf("%s isn't in your pokedex", pokemonToInspect)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Print("Stats:\n")
	for _, stats := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Print("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf("  -%s\n", types.Type.Name)
	}
	return nil
}