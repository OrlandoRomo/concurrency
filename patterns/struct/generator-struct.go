package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	msg     string
	waitFor chan bool
}

func main() {
	c := fanIn(returnChannel("Orlando"), returnChannel("Alma"))
loop:
	for {
		select {
		case msg := <-c:
			if (Message{}) == msg {
				break loop
			}
			msg.waitFor <- true
			fmt.Printf("You say %v\n", msg.msg)
		}
	}
}

// The generator pattern returns a a channel and this channel is generated in another function. A go routine is launched
// by this function instead of the main function.
func returnChannel(msg string) <-chan Message {
	c := make(chan Message)
	wait := make(chan bool)
	// Launching literally function
	go func() {
		for i := 0; i < 5; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-wait
		}
		c <- Message{}
	}()
	return c
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for _, input := range inputs {
		go func(i <-chan Message) {
			for {
				c <- <-i
			}
		}(input)
	}
	return c
}
