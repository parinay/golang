package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomNo int
}
type Result struct {
	job         Job
	sumofdigits int
}

// globals
// buffered channels

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number

	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)
	return sum
}

// Read from jobs channel and write to results channel
func worker(wg *sync.WaitGroup) {
	// range over bufferec channel jobs
	for job := range jobs {
		// digits() calculates the sum of digits &
		// updates the sum.
		// resulting structure is assigned to output
		output := Result{job, digits(job.randomNo)}

		// write to results channel the struct above
		results <- output
	}
	wg.Done()
}

// create the workerpool
func createWorkerPool(noOfWorkers int) {
	//WaitGroup to make sure all go routines finish before we exit
	var wg sync.WaitGroup

	for i := 0; i < noOfWorkers; i++ {
		// add to WaitGroup()
		wg.Add(1)
		// create worker go routine
		go worker(&wg)
	}
	// wait for all goroutines to finish
	wg.Wait()
	// close the result channel as nobody will be writing to these channes
	close(results)

}

// allocate jobs
func allocate(noOfJobs int) {
	// defer close of jobs buffered channel
	defer close(jobs)
	for i := 0; i < noOfJobs; i++ {
		// random number generator - generates peseudo random numbers
		// with max valye of 998

		randomNo := rand.Intn(999)
		// update the Job struct with values
		job := Job{i, randomNo}
		// write to jobs channel
		jobs <- job
	}
}

//print results
func result(done chan bool) {
	// range over results channle and print the sun of digits
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomNo, result.sumofdigits)
	}
	done <- true
}
func main() {

	// to bechmark time take to process number of jobs
	// record starttime and endtime
	startTime := time.Now()

	// 100 Jobs
	noOfJobs := 10
	go allocate(noOfJobs)

	// a boolean channel created
	// start result routine
	done := make(chan bool)
	go result(done)

	// 10 workers in workerpool to handle 100jobs
	noOfWorkers := 10
	//createWorkerPool()->worker() do the work of
	// calculating sum of digits and write to
	// results channel
	createWorkerPool(noOfWorkers)

	// write to done channel after worker() are done with
	// sum of digits and updating results channel
	// main() waits for on the done channel for all the results
	// to be printed
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total time taken ", diff.Seconds(), "seconds")
}
