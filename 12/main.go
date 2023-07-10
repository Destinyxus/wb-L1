package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	bunch := make(map[string]bool)

	for _, s := range sequence {
		bunch[s] = true
	}

	// Print the set
	for s := range bunch {
		fmt.Println(s)
	}
}
