package main

import "log"

type Pipe interface {
	Process(in chan int) chan int
}

type Square struct{}

func (sq *Square) Process(in chan int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

type Add struct{}

func (a *Add) Process(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			out <- i + i
		}
	}()
	return out
}

func main() {
	pipeline := NewPipeline(&Square{}, &Add{})

	go func() {

		for i := 0; i < 10; i++ {
			log.Printf("Sending: %v\n", i)
			pipeline.Enqueue(i)
		}
		log.Println("Closing pipeline")
		pipeline.Close()

	}()

	pipeline.Dequeue(func(i int) {
		log.Printf("Recevied: %v\n", i)
	})
}
