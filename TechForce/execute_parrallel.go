package main

import (
	"fmt"
	"sync"
)

func executeParallel(ch chan<- int, functions ...func() int) {
	var wg sync.WaitGroup

	for _, function := range functions {
		wg.Add(1)
		go func(f func() int) {
			defer wg.Done()
			ch <- f()
		}(function)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
}

func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}

func main() {
	expensiveFunction := func() int { return exampleFunction(200000000) }
	cheapFunction := func() int { return exampleFunction(10000000) }

	ch := make(chan int)

	// No need to pass sync.WaitGroup, handle it inside executeParallel
	go executeParallel(ch, expensiveFunction, cheapFunction)

	// Read results from the channel
	for result := range ch {
		fmt.Printf("Result: %d\n", result)
	}
}
