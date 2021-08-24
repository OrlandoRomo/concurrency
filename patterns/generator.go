package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(returnChannel("Al"), returnChannel("Prro"))
	// c := returnChannel("Al")
loop:
	for {
		select {
		case msg := <-c:
			if msg == "out" {
				fmt.Printf("Is the end %s\n", msg)
				break loop
			}
			fmt.Printf("You say %s\n", msg)
		}
	}
}

// The generator pattern returns a a channel and this channel is generated in another function. A go routine is launched
// by this function instead of the main function.
func returnChannel(msg string) <-chan string {
	c := make(chan string)
	// Launching literally function
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		c <- "out"
	}()
	return c
}
