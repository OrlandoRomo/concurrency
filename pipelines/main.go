package main

import (
	"fmt"
	"time"
)

type Data int

type Square int

type Result int

func main() {
	chanBufSize := 3

	raw := make(chan int, chanBufSize)

	dataChan := make(chan Data, chanBufSize)

	squareChan := make(chan Square, chanBufSize)

	stage3Result := make(chan Result, chanBufSize)

	stage2Result := make(chan Result, chanBufSize)

	go func() {
		for i := 0; i < 10; i++ {
			raw <- i
		}
	}()

	go Stage1(raw, dataChan, stage2Result)
	go Stage2(dataChan, squareChan, stage3Result, stage2Result)
	go Stage3(squareChan, stage3Result)

	time.Sleep(time.Second)

}

func Stage1(raw <-chan int, data chan<- Data, stage2Result <-chan Result) {
	fmt.Println("Starting stage 1")

	for {
		select {
		case rawData := <-raw:
			fmt.Printf("Stage 1: received raw from downstream: %d\n", rawData)
			data <- Data(rawData)
		case result := <-stage2Result:
			fmt.Printf("Stage 1: recevied result from upstream: %d\n", result)
		}
	}
}

func Stage2(data <-chan Data, squared chan<- Square, stage3Result <-chan Result, stage2Result chan<- Result) {
	fmt.Println("Starting stage 2")

	for {
		select {
		case d := <-data:
			fmt.Printf("Stage 2: received data from downstream: %d\n", d)
			sq := d * d
			squared <- Square(sq)
		case res := <-stage3Result:
			fmt.Printf("Stage 2: received result from upstream: %d\n", res)
			stage2Result <- res

		}
	}
}

func Stage3(squared <-chan Square, stage3Result chan<- Result) {
	fmt.Println("Starting stage 3")
	for {
		select {
		case d := <-squared:
			fmt.Printf("Stage 3: recevied data from downstream: %d\n", d)
			res := Result(3 * d)
			stage3Result <- res
			fmt.Printf("Stage 3: sent result data to downstream: %d\n", res)
		}
	}
}
