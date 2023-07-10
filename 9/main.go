package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	toProcess, processed := make(chan int), make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(1)

	// Launching one worker
	go func() {
		defer wg.Done()
		worker(ctx, toProcess, processed)
	}()

	go func() {
		wg.Wait()
		// Close channel when the worker is done
		// We close it to avoid block on range statement
		close(processed)
	}()

	go func() {
		defer close(toProcess)
		// Writing into channel
		for i := range arr {
			toProcess <- arr[i]
		}

	}()

	// Reading from channel with range
	for a := range processed {
		fmt.Println(a)
		// When "processed" channel is closed, it finishes
	}

}

func worker(ctx context.Context, toProc, proc chan int) {

	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-toProc:
			if !ok {
				return
			}
			proc <- d * 2
		}
	}
}
