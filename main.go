package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	next_offset int
}

func main() {
	//New scanner of terminal
	scanner := bufio.NewScanner(os.Stdin)

	//initialize config
	config := &config{
		next_offset: 0,
	}

	// Main loop of REPL
	for {
		fmt.Print("Pokedex> ")

		// Scan input and read to cleaner
		scanner.Scan()
		input := strings.Fields(strings.Trim(strings.ToLower(scanner.Text()), " "))

		// initialise args. If no input, continue. If more than one command, add to args
		var args []string
		if len(input) == 0 {
			continue
		} else if len(input) != 1 {
			args = input[1:]
		}

		// Get command. If valid command, execute
		cmd, ok := GetCommands()[input[0]]
		if ok {
			cmd.callback(config, args...)
		} else {
			fmt.Printf("Unknown command\n")
		}

	}
}

//
