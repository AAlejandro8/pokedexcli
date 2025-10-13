package main

import (
	"fmt"
	"errors"
)


func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.next)
	if err != nil {
		return err
	}
	
	for _, loc := range resp.Results {
		fmt.Printf(" - %v\n", loc.Name)
	}

	cfg.next = resp.Next
	cfg.prev = resp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prev == nil {
		return errors.New("you are already on the fist page")
	}
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.prev)
	if err != nil {
		return err
	}
	
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}
	cfg.next = resp.Next
	cfg.prev = resp.Previous
	return nil
}