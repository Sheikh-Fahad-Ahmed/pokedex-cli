package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sheikh-Fahad-Ahmed/pokedex-cli/internal/pokecache"
)

func GetMap(url string, config *Config, cache *pokecache.Cache) ([]Item, error) {
	var data []byte

	if cached, ok := cache.Get(url); ok {
		fmt.Println("Using cached data!...")
		data = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error Get request: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error io read: %w", err)
		}

		cache.Add(url, data)
		fmt.Println("data is Cached!..")
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("error unmarshal json data: %w", err)
	}

	return config.Results, nil

}
func GetEncounters(url string, config *PokemonEncountersResponse, cache *pokecache.Cache) ([]EncounterConfig, error) {
	var data []byte

	if cached, ok := cache.Get(url); ok {
		fmt.Println("Using cached data!")
		data = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("error get request: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error io read: %w", err)
		}

		cache.Add(url, data)
		fmt.Println("data is Cached!...")
	}

	if err := json.Unmarshal(data, config); err != nil {
		return nil, fmt.Errorf("error unmarshal json data: %w", err)
	}
	return config.PokemonEncountersConfig, nil
}

func GetPokeInfo(url string, cache * pokecache.Cache) (Pokemon, error) {
	var data []byte

	if cached, ok := cache.Get(url); ok {
		fmt.Println("Using cached data!")
		data = cached
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error get response: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error io read: %w", err)
		}

		cache.Add(url ,data)
		fmt.Println("data is Cached...")
	}
	
	var result Pokemon
	if err := json.Unmarshal(data, &result); err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshal json data: %w", err)
	}
	return result, nil 
}
