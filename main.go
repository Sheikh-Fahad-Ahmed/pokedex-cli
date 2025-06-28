package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var words []string
	word := ""
	for _, char := range text {
		if char == ' ' {
			if len(word) == 0 {
				continue
			}
			words = append(words, word)
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
