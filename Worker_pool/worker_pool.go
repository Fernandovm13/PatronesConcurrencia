package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("worker %d processing job %d\n", id, j)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	const numWorkers = 3
	jobs := make(chan int)
	var wg sync.WaitGroup

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
	fmt.Println("all jobs done")
}
