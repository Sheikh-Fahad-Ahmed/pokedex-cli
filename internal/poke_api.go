package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetMap() ([]Item, error) {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return nil, fmt.Errorf("error Get request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error io read: %w", err)
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error unmarshal json data: %w", err)
	}

	return config.Results, nil

}
