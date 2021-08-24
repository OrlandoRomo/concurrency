package main

func fanIn(inputs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, input := range inputs {
		go func(i <-chan string) {
			for {
				c <- <-i
			}
		}(input)
	}
	return c
}
