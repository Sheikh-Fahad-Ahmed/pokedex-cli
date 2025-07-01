package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetMap(url string, config *Config) ([]Item, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error Get request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error io read: %w", err)
	}

	if err := json.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error unmarshal json data: %w", err)
	}

	return config.Results, nil

}
