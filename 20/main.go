package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseSentence("snow dog sun"))
}

func reverseSentence(sentence string) string {
	words := strings.Fields(sentence) // Split the sentence into words

	reversedWords := make([]string, len(words))

	for i, j := 0, len(words)-1; i < len(words); i, j = i+1, j-1 {
		reversedWords[i] = words[j]
	}

	reversed := strings.Join(reversedWords, ",")

	return reversed
}
