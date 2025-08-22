package pokeapi

import (
	"errors"
	"io"
	"net/http"
)

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
