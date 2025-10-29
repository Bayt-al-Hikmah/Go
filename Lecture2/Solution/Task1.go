package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// Word
	word := "Developer"
	fmt.Printf("First byte: %c\n", word[0])
	fmt.Printf("Last byte: %c\n", word[len(word)-1])
	fmt.Println("Uppercase:", strings.ToUpper(word))

	// Phrase
	phrase := "I love Python"
	updatedPhrase := strings.Replace(phrase, "Python", "Go", 1)
	fmt.Println("Updated Phrase:", updatedPhrase)

	// Sentence
	sentence := "Learning Go is fun"
	containsGo := strings.Contains(sentence, "Go")
	fmt.Println("Contains 'Go':", containsGo)
	fmt.Println("Number of bytes:", len(sentence))
	fmt.Println("Number of runes:", utf8.RuneCountInString(sentence))
}