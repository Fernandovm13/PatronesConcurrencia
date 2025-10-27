package main

import (
	"fmt"
	"time"
)

func longComputation() <-chan int {
	res := make(chan int, 1)
	go func() {
		time.Sleep(500 * time.Millisecond)
		res <- 42
	}()
	return res
}

func main() {
	fut := longComputation()
	fmt.Println("doing other work...")
	result := <-fut
	fmt.Println("result:", result)
}
