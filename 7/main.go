package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	// In this task we will use mutex to prevent data race on reading into map
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	// Create a map to store some data
	data := make(map[string]int)

	numGoroutines := 10

	wg.Add(numGoroutines)

	// Perform concurrent writing
	for i := 0; i < numGoroutines; i++ {

		go func(index int) {
			defer wg.Done()
			mu.Lock()
			key := strconv.Itoa(index)
			value := index
			data[key] = value * 2
			mu.Unlock()

		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	for key, value := range data {
		fmt.Println(key, value)
	}
}
