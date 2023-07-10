package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var num int
	fmt.Println("Enter an amount of workers: ")
	fmt.Scan(&num)

	// Creating channel to receive an os Signals
	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Context for cancellation writing data if graceful shutdown occurs
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Creating an unbuffered channel, we also can make it buffered if we know an amount of data we want to send and read
	newData := make(chan interface{})

	// Launch workers, on each iteration we increment the waitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, newData)
		}()
	}
	// Implementation of graceful shutdown
	go func() {
		for {
			select {
			case sig := <-sigChan:
				fmt.Println("Graceful shutdown:", sig)
				// The program will cancel a context to make the workers to terminate their work
				cancel()

			}

		}

	}()

	// Creating a new goroutine which will write the data to a channel
	go func() {
		defer close(newData)
		for v := 0; v < 100000; v++ {
			// The reason we use here a select statement is when the graceful shutdown occurs, it terminates workers ,so we
			// have to stop writing the data to escape a deadlock.
			// When a graceful shutdown comes <-ctx.Done, so we close the channel and workers will end their job gracefully.
			select {
			case newData <- "some data":
			case <-ctx.Done():
				return
			}

		}

	}()

	// Here we wait the job termination of all workers.
	wg.Wait()
}

func worker(ctx context.Context, newData chan interface{}) {
	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-newData:
			// If channel is closed, terminate workers
			if !ok {
				return
			}
			fmt.Println(d)
		}

	}

}
