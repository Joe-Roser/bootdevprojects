package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cmdExit(_ *config, _ ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n\n")
	return errExit
}

func cmdHelp(_ *config, args ...string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func cmdMap(conf *config, args ...string) error {
	// Search cache and make request
	body, err := pokeapi.Get(conf.next_req, conf.cache)
	if err != nil {
		return err
	}

	// Parse response into locations struct
	var locations Locations

	err = json.Unmarshal([]byte(body), &locations)
	if err != nil {
		return err
	}

	// print all locations
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	// Handle updating conf
	if locations.Previous == nil {
		conf.prev_req = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		conf.prev_req = locations.Previous.(string)
	}
	if locations.Next == nil {
		conf.next_req = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		conf.next_req = locations.Next.(string)
	}

	return nil
}
func cmdMapb(conf *config, args ...string) error {
	// Search cache and make request
	body, err := pokeapi.Get(conf.prev_req, conf.cache)
	if err != nil {
		return err
	}

	// Parse response into locations struct
	var locations Locations

	err = json.Unmarshal([]byte(body), &locations)
	if err != nil {
		return err
	}

	// print all locations
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	// Handle updating conf
	if locations.Previous == nil {
		conf.prev_req = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		conf.prev_req = locations.Previous.(string)
	}
	if locations.Next == nil {
		conf.next_req = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		conf.next_req = locations.Next.(string)
	}

	return nil
}

func cmdExplore(conf *config, args ...string) error {
	if len(args) != 1 {
		return errInvalidInput
	}

	req := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", args[0])

	// Search cache and make request
	body, err := pokeapi.Get(req, conf.cache)
	if err != nil {
		return err
	}

	var location LocationReponse

	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Printf("Failed to unmarshall")
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("%v\n", pokemon.Pokemon.Name)
	}

	return nil
}

func cmdCatch(conf *config, args ...string) error {
	if len(args) != 1 {
		return errInvalidInput
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", args[0])
	body, err := pokeapi.Get(url, conf.cache)
	if err != nil {
		return err
	}

	var pokemon Pokemon

	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %s\n", args[0])

	wait := time.Duration((7+rand.Intn(10))*100) * time.Millisecond
	time.Sleep(wait)

	catch := 30 >= rand.Intn(pokemon.BaseExperience)
	if catch {
		fmt.Printf("%s was caught!!\n", strings.ToTitle(args[0]))
		conf.pokedex[args[0]] = pokemon
	} else {
		fmt.Printf("%s escaped!!\n", strings.ToTitle(args[0]))
	}

	return nil
}

//The commands dictionary

func GetCommands() map[string]cliCommand {
	// Return a map from keys to values
	return map[string]cliCommand{
		"catch": {
			name:        "catch [pokemon]",
			description: "Try catch a pokemon",
			callback:    cmdCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    cmdExit,
		},
		"explore": {
			name:        "explore [area]",
			description: "Name all pokemon in an area",
			callback:    cmdExplore,
		},
		"help": {
			name:        "help",
			description: "List all commands",
			callback:    cmdHelp,
		},
		"map": {
			name:        "map",
			description: "Show the map of the region",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Move back backwards",
			callback:    cmdMapb,
		},
		"test": {
			name:        "test",
			description: "test newest feature",
			callback:    building,
		},
	}
}

//
