package main

import (
	"fmt"
	"time"
)

func supervised(name string, stop <-chan struct{}) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("recovered:", r)
				}
			}()
			time.Sleep(200 * time.Millisecond)
			panic("simulated failure")
		}()
		select {
		case <-stop:
			fmt.Println("supervised", name, "stopping")
			return
		default:
			fmt.Println("restarting", name)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	stop := make(chan struct{})
	go supervised("worker1", stop)
	time.Sleep(700 * time.Millisecond)
	close(stop)
	time.Sleep(200 * time.Millisecond)
}
