package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan string)
	b := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		a <- "from a"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		b <- "from b"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-a:
			fmt.Println(msg)
		case msg := <-b:
			fmt.Println(msg)
		case <-time.After(150 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}
