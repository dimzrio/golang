package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var task = make(chan int)

func createTaks(noTaks int) {
	for i := 1; i <= noTaks; i++ {
		task <- i
	}
	// set channel task is done
	close(task)
}

func createWorker(noWorker int) {
	var workerGroup sync.WaitGroup

	for i := 1; i <= noWorker; i++ {
		workerGroup.Add(1)

		go func(args int, workerGroup *sync.WaitGroup) {

			// defer workerGroup.Done() --> defer is an fuction to call fuction after fuction call

			for valueTask := range task {
				fmt.Printf("Worker ke-%d, task-%d -> result random digit: %d\n", args, valueTask, rand.Intn(99))
				time.Sleep(1 * time.Second)
			}
			workerGroup.Done()

		}(i, &workerGroup)
	}
	workerGroup.Wait()
}
func main() {
	noTaks := 20
	noWorker := 5
	startTime := time.Now()

	go createTaks(noTaks)
	createWorker(noWorker)

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

/***
Another way with func read result

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var task = make(chan int)
var result = make(chan string)

func createTaks(noTaks int) {
	for i := 1; i <= noTaks; i++ {
		task <- i
	}
	// set channel task is done
	close(task)
}

func createWorker(noWorker int) {
	var workerGroup sync.WaitGroup

	for i := 1; i <= noWorker; i++ {
		workerGroup.Add(1)

		go func(args int, workerGroup *sync.WaitGroup) {

			// defer workerGroup.Done() --> defer is an fuction to call fuction after fuction call

			for valueTask := range task {
				result <- fmt.Sprintf("Worker ke-%d, task-%d -> result random digit: %d", args, valueTask, rand.Intn(99))
				time.Sleep(1 * time.Second)
			}
			workerGroup.Done()

		}(i, &workerGroup)
	}
	workerGroup.Wait()

	// set channel result is done
	close(result)
}

func getResult(done chan bool) {
	for message := range result {
		fmt.Println(message)
	}
	done <- true
}

func main() {
	noTaks := 20
	noWorker := 5
	startTime := time.Now()
	done := make(chan bool)

	go createTaks(noTaks)
	go getResult(done)
	createWorker(noWorker)

	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

***/
