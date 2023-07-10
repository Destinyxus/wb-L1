package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Creating a context which is responsible for termination program
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*20))

	defer cancel()

	wg := &sync.WaitGroup{}

	newData := make(chan interface{})

	wg.Add(1)
	// Creating a goroutine which writes the data
	go func() {
		defer wg.Done()
		for {
			select {
			case newData <- "some data":
			// if deadline occurs, we close the channel and terminate our worker
			case <-ctx.Done():
				close(newData)
				return
			}
		}
	}()

	go func() {
		for {
			select {
			// Reading some data from channel
			case d, ok := <-newData:
				// If it was closed earlier by the context, stop reading
				if !ok {
					return
				}
				fmt.Println(d)

			}

		}
	}()

	wg.Wait()
}
