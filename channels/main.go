package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // tell WaitGroup this goroutine is done when function exits

	for job := range jobs { // receive jobs from the jobs channel until it is closed
		fmt.Printf("Worker %d processing job %d\n", id, job)
		results <- job * 2 // send result into results channel
	}
}

func main() {

	jobs := make(chan int, 5)    // buffered channel — mailbox with 5 slots for jobs
	results := make(chan int, 5) // buffered channel — mailbox with 5 slots for results

	var wg sync.WaitGroup // create a WaitGroup to track goroutines

	// spawn 3 workers — each runs concurrently (goroutines)
	for i := 1; i <= 3; i++ {
		wg.Add(1)                        // tell WaitGroup we are starting 1 more goroutine
		go worker(i, jobs, results, &wg) // launch goroutine — runs in background
	}

	// send 5 jobs into the buffered channel — no blocking since buffer has space
	for j := 1; j <= 5; j++ {
		jobs <- j // drop job into mailbox
	}
	close(jobs) // close the channel — tells workers no more jobs are coming

	// wait in a separate goroutine — once all workers are done, close results
	go func() {
		wg.Wait()      // block here until all workers call wg.Done()
		close(results) // close results so we can range over it below
	}()

	// collect all results from the results channel
	for result := range results { // blocks until results channel is closed
		fmt.Println("Result:", result)
	}

}


// main creates 2 buffered channels  →  jobs mailbox and results mailbox

// 3 workers are launched as goroutines  →  all running concurrently in background

// 5 jobs are dropped into jobs channel  →  no blocking since buffer has 5 slots

// workers pick up jobs, process them, send results into results channel

// WaitGroup tracks all 3 workers
//   → each worker calls wg.Done() when finished
//   → wg.Wait() unblocks when all 3 are done
//   → results channel is closed

// main collects all results and prints them
