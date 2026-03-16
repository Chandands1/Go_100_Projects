package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {

	for job := range jobs {
		fmt.Println("Worker", id, "started job", job)

		time.Sleep(time.Second)

		result := job * 2

		fmt.Println("Worker", id, "finished job", job)

		results <- result
	}
}

func main() {

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 5; r++ {
		fmt.Println("Result:", <-results)
	}
}