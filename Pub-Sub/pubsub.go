package main

import (
	"fmt"
	"time"
)

type Broker struct {
	subs []chan string
}

func (b *Broker) Subscribe() chan string {
	ch := make(chan string, 2)
	b.subs = append(b.subs, ch)
	return ch
}

func (b *Broker) Publish(msg string) {
	for _, ch := range b.subs {
		select {
		case ch <- msg:
		default:
			go func(c chan string) { c <- msg }(ch)
		}
	}
}

func main() {
	b := &Broker{}
	s1 := b.Subscribe()
	s2 := b.Subscribe()

	b.Publish("hello subscribers")

	fmt.Println(<-s1)
	fmt.Println(<-s2)

	time.Sleep(100 * time.Millisecond)
}
