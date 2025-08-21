package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func CmdExit(_ *config, _ ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n\n")
	os.Exit(0)
	return nil
}

func CmdHelp(_ *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func CmdMap(_ *config, args ...string) error {
	return nil
}

func GetCommands() map[string]cliCommand {
	// Return a map from keys to values
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CmdExit,
		},
		"help": {
			name:        "help",
			description: "List all commands",
			callback:    CmdHelp,
		},
		"map": {
			name:        "map",
			description: "Show the map of the region",
			callback:    CmdMap,
		},
		"test": {
			name:        "test",
			description: "test newest feature",
			callback:    building,
		},
	}
}

//
