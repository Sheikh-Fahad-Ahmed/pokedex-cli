package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func init(){
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name: "help",
			description: "Display list of commands and the description",
			callback: commandHelp,
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
	fmt.Println("Usage: \n")
	for key, value := range commands {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	fmt.Println("")
	return nil
}
