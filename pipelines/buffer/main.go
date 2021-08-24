package main

import (
	"fmt"
	"sync"
)

// The main is the last stage of the pipeline, it's the consumer
func main() {
	in := gen(2, 3)
	c1 := sq(in)
	c2 := sq(in)

	done := make(chan struct{}, 2)
	out := merge(c1, c2)
	fmt.Println(<-out)

	done <- struct{}{}
	done <- struct{}{}
	// for n := range merge(c1, c2) {
	// 	fmt.Printf("SQ: %d\n", n)
	// }
}

// First stage of the pipeline generates an outbound channel
func gen(integers ...int) <-chan int {
	out := make(chan int, len(integers))
	for _, i := range integers {
		out <- i
	}
	close(out)
	return out
}

// Second stage of the pipeline
func sq(i <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range i {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Fan-In
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
