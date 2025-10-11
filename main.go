package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
	}
}

func main(){
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
		
		switch command.name {
		case "exit":
			commandExit()
		case "help":
			commandHelp()
		default:
			fmt.Println("Invalid command!")
		}
		
	}
}