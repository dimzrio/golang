package main

import (
	"fmt"
	"time"
)

var jobs = make(chan int, 10)
var result = make(chan string)

func process(id, job int) string {
	return fmt.Sprintf("Worker ke-%d Job ke-%d", id, job)
}

func consumer(id int) {
	for job := range jobs {
		result <- process(id, job)
		time.Sleep(time.Second * 1)
	}
}

func producer(task int) {
	for i := 1; i <= task; i++ {
		jobs <- i
	}
}

func main() {
	task := 10
	go producer(task)

	// Start consumer
	for i := 1; i <= 8; i++ {
		go consumer(i)
	}

	for i := 1; i <= 10; i++ {
		res := <-result
		fmt.Println(res)
	}

	close(jobs)
	close(result)

}
