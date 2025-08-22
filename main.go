package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"internal/pokecache"
)

// Errors
var ErrInvalidInput = errors.New("invalid input")
var ErrExit = errors.New("Exit")

// config type
type config struct {
	prev_req string
	next_req string
	cache    *pokecache.PokeCache
}

// Constants
const cache_reset = 30 * time.Second
const wait_after_exit = 500 * time.Millisecond

func main() {
	//New scanner of terminal
	scanner := bufio.NewScanner(os.Stdin)

	// Gets context to kill the cache if the program finishes
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
		time.Sleep(wait_after_exit)
	}()

	//initialize config
	config := &config{
		prev_req: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		next_req: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		cache:    pokecache.NewCache(ctx, cache_reset),
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
			fmt.Printf("Error: unknown command\n")
		}

		if errors.Is(err, ErrExit) {
			return
		} else if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		fmt.Println("")
	}
}

//
