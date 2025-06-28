package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands = map[string]cliCommand {
	"exit":{
		name : "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			line := scanner.Text()
			cleanedLine := cleanInput(line)
			
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
		words = append(words, word)
	}
	return words
}

// All Command Functions

func commandExit() {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
}