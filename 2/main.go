package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{2, 4, 6, 8, 10}

	for i := range arr {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			fmt.Println(math.Pow(float64(arr[v]), 2))
		}(i)
	}

	wg.Wait()
}
