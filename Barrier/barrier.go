package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ready *sync.WaitGroup) {
	defer ready.Done()
	fmt.Println("worker", id, "ready")
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Println("worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup
	n := 3
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all workers reached barrier — continue")
}
