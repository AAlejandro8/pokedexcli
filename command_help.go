package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("%v : %v\n", cmd.name, cmd.description)
	}
	return nil
}