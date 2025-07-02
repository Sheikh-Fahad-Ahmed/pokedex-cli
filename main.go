package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Sheikh-Fahad-Ahmed/pokedex-cli/internal/api"
	"github.com/Sheikh-Fahad-Ahmed/pokedex-cli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*api.Config, *pokecache.Cache) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display list of commands and the description",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display list of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map back",
			description: "Go to the previous list of the map",
			callback:    commandMapBack,
		},
	}
}

func main() {
	config := &api.Config{}
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(30 * time.Second)


	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			cleanedLine := cleanInput(line)
			val, ok := commands[cleanedLine[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				err := val.callback(config, cache)
				if err != nil {
					fmt.Println("error running function ", err)
				}
			}
		}

	}
}

func cleanInput(text string) []string {
	var words []string
	word := ""
	for _, char := range text {
		if char == ' ' {
			if len(word) == 0 {
				continue
			}
			words = append(words, strings.ToLower(word))
			word = ""
			continue
		}
		word += string(char)
	}
	if len(word) != 0 {
		words = append(words, strings.ToLower(word))
	}
	return words
}

// All Command Functions

func commandExit(c *api.Config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp(c *api.Config, cache *pokecache.Cache) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage: \n")
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(c *api.Config, cache *pokecache.Cache) error {
	var result []api.Item
	var err error
	url := "https://pokeapi.co/api/v2/location-area"

	if c.Count == 0 {
		result, err = api.GetMap(url, c, cache)
	} else {
		result, err = api.GetMap(*c.Next, c, cache)
	}

	if err != nil {
		return fmt.Errorf("the error: %w", err)
	}
	c.Count += 1

	for _, item := range result {
		fmt.Println(item.Name)
	}

	return nil
}

func commandMapBack(c *api.Config, cache *pokecache.Cache) error {
	if c.Previous == nil {
		fmt.Println("You are on the first page...")
		c.Count = 0
		return nil
	}

	result, err := api.GetMap(*c.Previous, c, cache)
	if err != nil {
		return fmt.Errorf("the error: %w", err)
	}
	for _, item := range result {
		fmt.Println(item.Name)
	}

	c.Count -= 1
	return nil
}

