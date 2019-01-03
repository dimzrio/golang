package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func consume(data <-chan int, done chan<- bool) {
	file, err := os.Create("example.txt")

	if err != nil {
		log.Fatal(err)
		done <- false
	}

	for write := range data {
		_, errWrite := fmt.Fprintln(file, write)

		if errWrite != nil {
			log.Fatal(err)
			file.Close()
			done <- false
			return
		}
	}
	file.Close()
	done <- true
}

func produce(data chan<- int, workerGroup *sync.WaitGroup) {
	number := rand.Intn(999)
	data <- number
	workerGroup.Done()
}

func main() {
	data := make(chan int)
	done := make(chan bool)

	var workerGroup sync.WaitGroup

	for i := 1; i <= 100; i++ {
		workerGroup.Add(1)
		go produce(data, &workerGroup)
	}

	go func() {
		workerGroup.Wait()
		close(data)
	}()

	go consume(data, done)

	status := <-done

	if status != true {
		fmt.Println("File writing failed")
		return
	}
	fmt.Println("File writing successfully")

}
