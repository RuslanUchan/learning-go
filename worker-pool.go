package main

import (
	"fmt"
	"time"
)

/* 	IMPORTANT CONCEPT TO GRASP
chan<- == setting values to the channel. Setter
<-chan == get values from the channel. Getter
*/

func main() {
	fmt.Println("Worker Pool")

	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// worker invocations
	// go consumer(1, jobs, results)
	// go consumer(2, jobs, results)
	for i := 1; i <= 4; i++ {
		go consumer(i, jobs, results)
	}

	go producer(jobs, 10) // concurrent running producer

	for i := 1; i <= 10; i++ {
		res := <-results
		fmt.Println("Results of", res)
	}

	fmt.Println("==================")
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)
}

func fakeHTTPRequest(x int) int {
	// fake jobs for illustrating worker pool
	return x * 10
}

func producer(jobs chan<- int, size int) {
	/* 	producer of jobs for workers
	write only channel producer indicated by (jobs chan<- int)
	manually specify channel size (size int)
	size indicate number of jobs to loop through in the channel
	*/

	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, results chan<- int) {
	/*	worker function. consume jobs
		get jobs from channel
		set result values to the channel
	*/

	for job := range jobs {
		fmt.Println("Worker/Consumer of", id)

		time.Sleep(time.Second * 2)

		results <- fakeHTTPRequest(job)
	}
}
