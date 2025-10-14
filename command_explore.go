package main

import "fmt"

func callbackExplore(cfg *config) error {
	resp, err := cfg.pokeapiClient.ExploreLocation(cfg.location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", cfg.location)
	fmt.Println("Found Pokemon:")
	for _, encounters := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}
	return nil
}