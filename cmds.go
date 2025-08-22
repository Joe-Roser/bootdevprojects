package main

import (
	"encoding/json"
	"fmt"

	"internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cmdExit(_ *config, _ ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n\n")
	return ErrExit
}

func cmdHelp(_ *config, args ...string) error {
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

type LocationReponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func cmdMap(conf *config, args ...string) error {
	// Check cache
	body, ok := conf.cache.Get(conf.next_req)

	//if no hit, make requset
	var err error
	if !ok {
		body, err = pokeapi.MakeRequest(conf.next_req)
		if err != nil {
			return err
		}
		conf.cache.Add(conf.next_req, body)
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
	// Check Cache
	body, ok := conf.cache.Get(conf.prev_req)

	//if no hit, make requset
	var err error
	if !ok {
		body, err = pokeapi.MakeRequest(conf.prev_req)
		if err != nil {
			return err
		}
		conf.cache.Add(conf.prev_req, body)
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
		return ErrInvalidInput
	}

	req := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", args[0])

	// Check cache
	body, ok := conf.cache.Get(req)

	//if no hit, make requset
	var err error
	if !ok {
		body, err = pokeapi.MakeRequest(req)
		if err != nil {
			fmt.Printf("Http request failed!!\n")
			return err
		}
		conf.cache.Add(conf.next_req, body)
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

//The commands dictionary

func GetCommands() map[string]cliCommand {
	// Return a map from keys to values
	return map[string]cliCommand{
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
