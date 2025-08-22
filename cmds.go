package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

type Locations struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CmdMap(conf *config, args ...string) error {
	//get request
	res, err := http.Get(conf.next_req)
	if err != nil {
		return err
	}

	// Read response and check its valid
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New("Api request failed")
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
		conf.prev_req = "https://pokeapi.co/api/v2/location-area/"
	} else {
		conf.prev_req = locations.Previous.(string)
	}
	if locations.Next == nil {
		conf.next_req = "https://pokeapi.co/api/v2/location-area/"
	} else {
		conf.next_req = locations.Next.(string)
	}

	return nil
}
func CmdMapb(conf *config, args ...string) error {
	//get request
	res, err := http.Get(conf.prev_req)
	if err != nil {
		return err
	}

	// Read response and check its valid
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New("Api request failed")
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
		"mapb": {
			name:        "mapb",
			description: "Move back backwards",
			callback:    CmdMapb,
		},
		"test": {
			name:        "test",
			description: "test newest feature",
			callback:    building,
		},
	}
}

//
