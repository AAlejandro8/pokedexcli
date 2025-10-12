package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AAlejandro8/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next *string
	prev *string
}

type cliCommand struct {
	name string
	description string
	callback func(*config) error
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:		"exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: 		"help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:		"map",
			description: "Displays locations",
			callback: callbackMap,
		},
		"mapb": {
			name:  		"mapb",
			description: "Displays the last 20 locations",
			callback: callbackMapb,	
		},
	}
}

func startRepl(cfg *config){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}
		
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command!")
			continue
		}
		
		if err := command.callback(cfg); err != nil {
			fmt.Println("error", err)
		}
		
	}
}