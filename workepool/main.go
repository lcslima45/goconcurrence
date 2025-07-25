package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	id  int
	num int
}

type Result struct {
	job   Job
	total int
}

var (
	jobs    = make(chan Job, 10)
	results = make(chan Result, 10)
)

func sum(number int) (total int) {
	no := number
	for no != 0 {
		digit := no % 10
		total += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, sum(job.num)}
		results <- output
	}
	wg.Done()
}

func createWorkePool() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocateJob() {
	for i := 0; i < 300; i++ {
		num := rand.Intn(999)
		job := Job{i, num}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, numero %d soma digitos %d\n",
			result.job.id,
			result.job.num,
			result.total,
		)
		time.Sleep(5 * time.Second)
	}
	done <- true
}

func main() {
	go func() {
		for {
			fmt.Println("Goroutines:", runtime.NumGoroutine())
			time.Sleep(2 * time.Second)
		}
	}()
	go allocateJob()
	done := make(chan bool)
	go result(done)
	createWorkePool()

	<-done
}
