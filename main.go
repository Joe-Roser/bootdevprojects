package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	prev_req string
	next_req string
}

func main() {
	//New scanner of terminal
	scanner := bufio.NewScanner(os.Stdin)

	//initialize config
	config := &config{
		prev_req: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		next_req: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	}

	// Main loop of REPL
	for {
		fmt.Print("Pokedex> ")

		// Scan input and read to cleaner
		scanner.Scan()
		input := strings.Fields(strings.Trim(strings.ToLower(scanner.Text()), " "))
		fmt.Println("")

		// initialise args. If no input, continue. If more than one command, add to args
		var args []string
		if len(input) == 0 {
			continue
		} else if len(input) != 1 {
			args = input[1:]
		}

		// Get command. If valid command, execute
		cmd, ok := GetCommands()[input[0]]
		var err error
		if ok {
			err = cmd.callback(config, args...)
		} else {
			fmt.Printf("Unknown command\n")
		}
		if err != nil {
			fmt.Printf("Error: %v. Shutting down\n", err)
			os.Exit(0)
		}
		fmt.Println("")

	}
}

//
