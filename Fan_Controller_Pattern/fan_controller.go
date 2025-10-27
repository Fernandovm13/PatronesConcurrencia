package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, quit <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case j, ok := <-jobs:
			if !ok {
				fmt.Println("worker", id, "jobs closed")
				return
			}
			fmt.Printf("worker %d processed job %d\n", id, j)
			time.Sleep(150 * time.Millisecond)
		case <-quit:
			fmt.Println("worker", id, "stopping (quit)")
			return
		}
	}
}

func main() {
	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	minWorkers := 1
	maxWorkers := 6
	currentWorkers := 0
	var quitChans []chan struct{}
	var mu sync.Mutex

	startWorker := func() {
		mu.Lock()
		defer mu.Unlock()
		currentWorkers++
		q := make(chan struct{})
		quitChans = append(quitChans, q)
		wg.Add(1)
		id := currentWorkers
		go worker(id, jobs, q, &wg)
		fmt.Println("started worker", id, "total:", currentWorkers)
	}

	for i := 0; i < minWorkers; i++ {
		startWorker()
	}

	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		for range ticker.C {
			lenJobs := len(jobs)
			mu.Lock()
			cw := currentWorkers
			mu.Unlock()
			if lenJobs > 10 && cw < maxWorkers {
				startWorker()
			}
		}
	}()

	go func() {
		for j := 1; j <= 40; j++ {
			jobs <- j
			time.Sleep(30 * time.Millisecond)
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("all workers done")
}
