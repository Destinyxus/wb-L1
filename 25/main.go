package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Before sleep")

	// Call the custom sleep function for 2 seconds
	sleep(2 * time.Second)

	fmt.Println("After sleep")
}
