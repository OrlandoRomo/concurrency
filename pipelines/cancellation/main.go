package main

import "fmt"

// The main is the last stage of the pipeline, it's the consumer
func main() {
	for n := range sq(sq(gen(1, 2, 3, 4, 5, 6, 7))) {
		fmt.Printf("SQ: %d\n", n)
	}
}

// First stage of the pipeline generates an outbound channel
func gen(integers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range integers {
			out <- i
		}
		close(out)
	}()
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
