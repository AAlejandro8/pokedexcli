package main

import "fmt"

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Location not provided")
	}
	location := args[0]
	resp, err := cfg.pokeapiClient.ExploreLocation(location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounters := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}
	return nil
}