package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Analytics struct {
	mu    sync.RWMutex
	stats map[string]int
}

func (a *Analytics) Increment(key string) {
	a.mu.Lock() // Exclusive lock for writing
	a.stats[key]++
	a.mu.Unlock()
}

func (a *Analytics) Get(key string) int {
	a.mu.RLock() // shared for reading
	defer a.mu.RUnlock()
	return a.stats[key]
}

func main() {
	data := &Analytics{stats: make(map[string]int)}
	//create a context that cancels after 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done(): //signal received to stop
				fmt.Println("Worker stopping: Context cancelled")
				return
			default:
				data.Increment("hits")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Printf("Final Hits: %d\n", data.Get("hits"))
}
