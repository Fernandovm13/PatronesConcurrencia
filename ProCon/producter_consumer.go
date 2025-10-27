package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i
		fmt.Println("produced", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println("consumed", v)
		time.Sleep(150 * time.Millisecond)
	}
	done <- struct{}{}
}

func main() {
	ch := make(chan int, 5)
	done := make(chan struct{})
	go producer(ch, 10)
	go consumer(ch, done)
	<-done
}
