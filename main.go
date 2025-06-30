package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/Sheikh-Fahad-Ahmed/pokedex-cli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			callback:    getMapList,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			cleanedLine := cleanInput(line)
			val, ok := commands[cleanedLine[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				err := val.callback()
				if err != nil {
					fmt.Println("Error running the function", err)
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Printf("Usage: \n")
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	fmt.Println("")
	return nil
}

func getMapList() error {
	result, err := internal.GetMap()
	if err != nil {
		return fmt.Errorf("the error: %w", err)
	}
	for _, item := range result {
		fmt.Println(item.Name)
	} 
	return nil
}
