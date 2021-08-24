package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	numbers := GenerateNumbers(1e7)

	t := time.Now()
	_ = Add(numbers)
	fmt.Printf("Sequential Add took: %s\n", time.Since(t))

	t = time.Now()
	_ = AddConcurrent(numbers)
	fmt.Printf("Concurrent Add took: %s\n", time.Since(t))

}

func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add returns the sum of numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

func AddConcurrent(numbers []int) int64 {
	// Use all cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	numOfCors := runtime.NumCPU()
	fmt.Printf("Print this alv %v\n", 348140/runtime.NumCPU())

	var sum int64

	max := len(numbers)

	sizeOfParts := max / numOfCors

	var wg sync.WaitGroup

	for i := 0; i < numOfCors; i++ {
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := numbers[start:end]

		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()

			var partSum int64
			for _, n := range nums {
				partSum += int64(n)
			}

			atomic.AddInt64(&sum, partSum)
		}(part)
	}

	wg.Wait()
	return sum
}
