package main

import (
	"fmt"
	"sync"
)

var x = 0

func main() {
	/* race conditions */
	// var result sync.WaitGroup

	// for i := 0; i < 100; i++ {
	// 	result.Add(1)

	// 	go func(result *sync.WaitGroup) {
	// 		defer result.Done()
	// 		x++
	// 	}(&result)
	// }
	// result.Wait()
	// fmt.Println("Race Conditions X: ", x)

	/* Solution */
	var solution sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		solution.Add(1)

		go func(solution *sync.WaitGroup, m *sync.Mutex) {
			defer solution.Done()

			m.Lock()
			x++
			m.Unlock()

		}(&solution, &mutex)
	}
	solution.Wait()
	fmt.Println("Race Solutions X: ", x)
}
