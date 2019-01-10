package main

import (
	"fmt"
	"time"
)

// process fucntion
func process(value int) int {
	return value * 10
}

// producer func with arg channel jobs only read channel
func producer(jobs chan<- int, size int) {
	for i := 1; i <= size; i++ {
		jobs <- i
	}
	close(jobs)
}

func consumer(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("Mulai job ke : ", id)
		time.Sleep(time.Second * 1)
		result <- process(job)
	}

}

func main() {
	jobs := make(chan int, 10)
	result := make(chan int, 10)

	go producer(jobs, 10)
	go consumer(1, jobs, result)
	go consumer(2, jobs, result)
	go consumer(3, jobs, result)

	for i := 1; i <= 10; i++ {
		res := <-result
		fmt.Println("Hasil: ", res)
	}

}
