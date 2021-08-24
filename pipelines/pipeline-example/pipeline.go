package main

type Pipeline struct {
	head chan int
	tail chan int
}

func (p *Pipeline) Enqueue(item int) {
	p.head <- item
}

func (p *Pipeline) Dequeue(handler func(int)) {
	for i := range p.tail {
		handler(i)
	}
}

func (p *Pipeline) Close() {
	close(p.head)
}

func NewPipeline(pipes ...Pipe) *Pipeline {
	head := make(chan int)
	var next_chan chan int
	for _, pipe := range pipes {
		if next_chan == nil {
			next_chan = pipe.Process(head)
		} else {
			next_chan = pipe.Process(next_chan)
		}
	}
	return &Pipeline{head: head, tail: next_chan}
}
