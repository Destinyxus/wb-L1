package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// First approach with wait group
	/*ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		for {
			select {
			case d, ok := <-ch:
				if !ok {
					wg.Done()
					return
				}
				fmt.Println(d)

			}
		}

	}()

	for i := 0; i < 1000; i++ {

		ch <- i

	}
	close(ch)
	wg.Wait()
	*/

	// Second approach with context

	/* ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- 123:
			}

		}

	}()
	for {
		select {
		case <-ctx.Done():
			return
		case d, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(d)
		}
	}
	*/

	// Third approach with a channel struct

	/*doneChan := make(chan struct{})
	ch := make(chan int)

	go func() {
		defer close(ch)
		for {
			select {
			case <-doneChan:
				return
			case ch <- 123:
			}
		}

	}()

	for i := 0; i < 5; i++ {
		select {
		case <-doneChan:
			return
		case d, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(d)
		}

	}
	close(doneChan)*/

	var wg sync.WaitGroup
	ch := make(chan int)
	terminate := make(chan struct{})
	cond := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-terminate:
					return
				case ch <- id:

				}
			}
		}(i)
	}

	go func() {
		for {
			select {
			case d, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(d)
			}

		}
	}()

	time.Sleep(time.Second)
	close(terminate)

	// Notify all goroutines to check for termination
	cond.Broadcast()

	// Wait for all goroutines to finish
	wg.Wait()
}
