package pokeapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"internal/pokecache"
)

func Get[T any](url string, c *pokecache.PokeCache) (T, error) {
	var zero T

	// if value is in cache, return it
	if value, ok := c.Get(url); ok {
		if value, ok := value.(T); ok {
			return value, nil
		}
		return zero, errors.New("cached value has different type")
	}

	//Else, make request
	res, err := http.Get(url)
	if err != nil {
		return zero, err
	}
	if res.StatusCode > 299 {
		return zero, errors.New("Api request failed")
	}

	// Make a decoder from the request
	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	// Decode into the type they requested
	var parsed T
	if err := decoder.Decode(&parsed); err != nil {
		return zero, err
	}

	// Add to the cache
	c.Add(url, parsed)

	// Return
	return parsed, nil
}

//
