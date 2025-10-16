package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/AAlejandro8/pokedexcli/internal/pokeapi"
)
type Pokedex map[string]pokeapi.Pokemon

type config struct {
	pokeapiClient pokeapi.Client
	pokedex Pokedex
	next *string
	prev *string
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
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
		"explore": {
			name: 		"explore <location>",
			description: "Displays all pokemon in the area",
			callback: callbackExplore,
		},
		"catch" : {
			name: 		"catch",
			description: "Chance to catch a pokemon",
			callback: callbackCatch,
		},
		"inspect" : {
			name:		"inspect",
			description: "inspect the pokemon in your pokedex",
			callback: callbackInspect,
		},
		"pokedex" : {
			name:		"pokedex",
			description: "Display all your pokemon in your pokedex",
			callback: callBackPokedex,
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
		
		// extract info 
		commandName := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 { 
			args = cleaned[1:]
		}

		// get the commands
		availableCommands := getCommands()

		// does the command we call even exist?
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command!")
			continue
		}

		// make the call back with the info needed
		if err := command.callback(cfg, args...); err != nil {
			fmt.Println("error", err)
		}
		
	}
}