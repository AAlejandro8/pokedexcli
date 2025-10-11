package main

import (
	"fmt"
	"github.com/AAlejandro8/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	url := ""
	if cfg.next != nil {
		url = *cfg.next
	}
	list, err := pokeapi.getLocationAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range list.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = list.Next
	cfg.prev = list.Previous
	return nil
}


func commandMapB(cfg *config) error {
	if cfg.next == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	list, err := pokeapi.getLocationAreas(url)
	if err != nil {
		return err
	}

	for _, loc := range list.Results {
		fmt.Println(loc.Name)
	}

	cfg.next = list.Next
	cfg.prev = list.Previous
	return nil
}
