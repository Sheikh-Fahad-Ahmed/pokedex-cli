package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	"github.com/Sheikh-Fahad-Ahmed/pokedex-cli/internal/api"
	"github.com/Sheikh-Fahad-Ahmed/pokedex-cli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(any, string, *pokecache.Cache) error
	configKind  string
}

func configHandler[T any](fn func(*T, string, *pokecache.Cache) error) func(any, string, *pokecache.Cache) error {
	return func(config any, param string, cache *pokecache.Cache) error {
		typedConfig, ok := config.(*T)
		if !ok {
			return fmt.Errorf("expected config type %T", new(T))
		}

		return fn(typedConfig, param, cache)
	}
}

var pokedex = make(map[string]api.Pokemon)
var commands map[string]cliCommand

func init() {

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    configHandler(commandExit),
			configKind:  "mapConfig",
		},
		"help": {
			name:        "help",
			description: "Display list of commands and the description",
			callback:    configHandler(commandHelp),
			configKind:  "mapConfig",
		},
		"map": {
			name:        "map",
			description: "Display list of locations",
			callback:    configHandler(commandMap),
			configKind:  "mapConfig",
		},
		"mapb": {
			name:        "map back",
			description: "Go to the previous list of the map",
			callback:    configHandler(commandMapBack),
			configKind:  "mapConfig",
		},
		"explore": {
			name:        "explore",
			description: "Displays the list of all pokemon in that specific area",
			callback:    configHandler(commandExplore),
			configKind:  "encounterConfig",
		},
		"catch": {
			name:        "catch",
			description: "Throws a Pokeball to catch the specified Pokemon",
			callback:    configHandler(commandCatch),
			configKind:  "pokemonConfig",
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the stats of a Pokemon in your Pokedex",
			callback:    configHandler(commandInspect),
			configKind:  "pokemonConfig",
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the Pokemons in your Pokedex",
			callback:    configHandler(commandPokedex),
			configKind:  "pokemonConfig",
		},
	}
}

func main() {

	configs := map[string]any{
		"mapConfig":       &api.Config{},
		"encounterConfig": &api.PokemonEncountersResponse{},
		"pokemonConfig":   &api.Pokemon{},
	}

	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(30 * time.Second)

	for {
		fmt.Printf("\nPokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			cleanedLine := cleanInput(line)
			if len(cleanedLine) == 0 {
				continue
			}
			val, ok := commands[cleanedLine[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			config, ok := configs[val.configKind]
			if !ok {
				fmt.Println("unknown config type")
				continue
			}

			param := ""
			if len(cleanedLine) > 1 {
				param = cleanedLine[1]
			}

			err := val.callback(config, param, cache)
			if err != nil {
				fmt.Println("the error: ", err)
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

func checkCatch(num int) bool {
	// percent := 0.45
	min := num / 2
	max := int(float64(num) * 1.5)
	return rand.IntN(max-min+1)+min >= num
}

// All Command Functions

func commandExit(c *api.Config, param string, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp(c *api.Config, param string, cache *pokecache.Cache) error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(c *api.Config, param string, cache *pokecache.Cache) error {
	var result []api.Item
	var err error
	url := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

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

func commandMapBack(c *api.Config, param string, cache *pokecache.Cache) error {
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

func commandExplore(e *api.PokemonEncountersResponse, param string, cache *pokecache.Cache) error {
	if param == "" {
		fmt.Println("Please specify a Location from the list given by map command....")
		return nil
	}
	fmt.Println(param)
	url := "https://pokeapi.co/api/v2/location-area"
	fullURL := fmt.Sprintf("%s/%s", url, param)
	results, err := api.GetEncounters(fullURL, e, cache)
	if err != nil {
		return fmt.Errorf("the error : %w", err)
	}
	fmt.Println("Exploring " + param + "...")
	fmt.Println("Found Pokemon:")
	for _, item := range results {
		fmt.Printf(" - %s\n", item.PokemonEncounter.Name)
	}
	return nil
}

func commandCatch(p *api.Pokemon, param string, cache *pokecache.Cache) error {
	if param == "" {
		fmt.Println("You need to specify a Pokemon...")
		return nil
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", param)
	fmt.Printf("Throwing a Pokeball at %s...\n", param)
	result, err := api.GetPokeInfo(url, cache)
	if err != nil {
		return fmt.Errorf("catch command error: %w", err)
	}

	if checkCatch(result.BaseEXP) {
		fmt.Printf("%s was caught!\n", param)
		pokedex[param] = result
	} else {
		fmt.Printf("%s escaped!\n", param)
	}
	return nil
}

func commandInspect(p *api.Pokemon, param string, cache *pokecache.Cache) error {
	if param == "" {
		fmt.Println("Please specify a Pokemon...")
		return nil
	}
	pokemon, ok := pokedex[param]
	if !ok {
		fmt.Println("You have not caught this pokemon")
	} else {
		fmt.Printf("Name: %s\nHeight: %d\nWeight:%d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.StatInfo.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, types := range pokemon.Types {
			fmt.Printf(" - %s\n", types.Type.Name)
		}
	}
	return nil
}

func commandPokedex(p *api.Pokemon, param string, cache *pokecache.Cache) error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty...\nTry catching some Pokemons!")
		return nil
	}
	fmt.Println("Your Pokemons:")
	for pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
