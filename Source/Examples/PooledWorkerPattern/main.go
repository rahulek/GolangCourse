package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	//Number of jobs
	numJobs = 10
)

func main() {
	fmt.Println("----- PooledWorker Pattern : START ------")
	fmt.Println()

	//Buffered Channels for the Jobs and their results
	jobsCh := make(chan int, numJobs)
	resultsCh := make(chan int, numJobs)

	//Start few workers: less workers(5)
	//than the number of jobs (10)
	for i := 0; i < 5; i++ {
		go worker(i, jobsCh, resultsCh)
	}

	//Create work jobs
	for j := 0; j < numJobs; j++ {
		jobsCh <- j + 1
	}
	//All jobs finished
	close(jobsCh)

	//Get all the results
	for r := 0; r < numJobs; r++ {
		fmt.Printf("\tJobs result: %d\n", <-resultsCh)
	}
	//Finished fetching the results
	close(resultsCh)

	fmt.Println()
	fmt.Println("----- PooledWorker Pattern : END ------")
}

// worker goroutine that waits for the jobs
// and spawns a new goroutine for each new job
// received
func worker(id int, jobsCh <-chan int, resultsCh chan<- int) {
	//Sync all the worker goroutines on a waitGroup
	wg := sync.WaitGroup{}

	//Pick all the jobs
	for job := range jobsCh {
		//Track a goroutine that will do the actual work
		wg.Add(1)
		//spawn a goroutine
		go func(jobId int) {
			defer wg.Done()

			//Simulate the work being done
			fmt.Printf("---> Started Job: %d\n", jobId)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("---> Job %d finished.\n", jobId)

			//Fake the result
			resultsCh <- jobId * 10
		}(job)
	}
	//Wait for all the worker goroutines to finish
	wg.Wait()
}
