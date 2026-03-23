package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment(wg *sync.WaitGroup) {
	defer wg.Done() // Signal we are finished when function ends
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	counter := SafeCounter{}
	for i := range 1000 {
		fmt.Printf("go routine %d created\n", i)
		wg.Add(1)
		go counter.Increment(&wg) // start 1000 independent tasks
	}
	wg.Wait() // wait for all 1000 to finish
	fmt.Println("Final counter:", counter.value)

	fmt.Println("--------------The Channel Approach------------")
	//1. Create a channel to send increment requests
	//We use a simple bool channel because the signal is all we need
	incrementChan := make(chan bool)

	//2. State-management
	//Only part that touches count
	count := 0

	go func() {
		for range incrementChan {
			count++
		}
	}()

	//Launch 1000 workers

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementChan <- true // send signal to increment
		}()
	}
	wg.Wait()
	close(incrementChan) // closing stops "state manager" loop

	fmt.Println("Final counter", count)
}
