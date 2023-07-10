package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("abAB"))
}

func isUnique(char string) bool {

	for i, v := range []rune(char) {

		for j, k := range []rune(char) {
			if i != j && strings.ToLower(string(v)) == strings.ToLower(string(k)) {
				return false
			}

		}
	}
	return true

	/*str = strings.ToLower(str)

	// Create a map to store encountered characters
	encountered := make(map[rune]bool)

	// Iterate over each character in the string
	for _, char := range str {
		// If the character is already encountered, it's not unique
		if encountered[char] {
			return false
		}

		encountered[char] = true
	}

	return true*/
}
