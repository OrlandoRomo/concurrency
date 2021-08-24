package main

import "fmt"

type Result struct {
	Num     int
	Display string
}

func fizzbuzz(out chan Result) {
	for num := 1; ; num++ {
		switch {
		case num%3 == 0 && num%5 == 0:
			out <- Result{num, "FizzBuzz"}
		case num%3 == 0:
			out <- Result{num, "Fizz"}
		case num%5 == 0:
			out <- Result{num, "Buzz"}
		default:
			out <- Result{num, fmt.Sprintf("%d", num)}
		}
	}
}

func main() {

	out := make(chan Result)
	go fizzbuzz(out)

	for res := range out {
		if res.Num >= 100 {
			break
		}
		fmt.Println(res.Display)
	}
}
