package main

import (
	"fmt"
	"sync"
)

// The main is the last stage of the pipeline, it's the consumer
func main() {
	done := make(chan struct{})
	defer close(done)

	in := gen(done, 2, 3)

	c1 := sq(done, in)
	c2 := sq(done, in)

	out := merge(done, c1, c2)
	fmt.Println(<-out)
}

// First stage of the pipeline generates an outbound channel
func gen(done <-chan struct{}, integers ...int) <-chan int {
	out := make(chan int, len(integers))
	for _, i := range integers {
		out <- i
	}
	close(out)
	return out
}

// Second stage of the pipeline
func sq(done <-chan struct{}, i <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range i {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// Fan-In
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
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
