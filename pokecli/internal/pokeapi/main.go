package pokeapi

import (
	"errors"
	"io"
	"net/http"

	"internal/pokecache"
)

func Get(url string, c *pokecache.PokeCache) ([]byte, error) {
	body, ok := c.Get(url)

	// Return on cache hit
	if ok {
		return body, nil
	}

	// else, make request, check errors and return
	body, err := MakeRequest(url)
	if err != nil {
		return nil, err
	}

	// remember to add the request to the cache
	c.Add(url, body)
	return body, nil
}

func MakeRequest(url string) ([]byte, error) {
	//get request
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	// Read response and check its valid
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []byte{}, err
	}
	if res.StatusCode > 299 {
		return []byte{}, errors.New("Api request failed")
	}

	return body, err
}
