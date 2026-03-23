package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Ping Pong Game

func main() {
	//1. Creating a channel
	table := make(chan string)
	//2. Launch a player

	go player("Ping", table)
	go player("Pong", table)

	//3. Serve the ball
	table <- "ball"

	//4. Let them play for 2 seconds
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Game Over: Main thread exiting...")

	// rand.Seed(time.Now().UnixNano()) // for randomness
	results := make(chan string)

	go fetchPrice("Binance", results)
	go fetchPrice("Coinbase", results)
	go fetchPrice("Kraken", results)

	select {
	case res := <-results:
		fmt.Println("Winner: ", res)
	case <-time.After(500 * time.Millisecond):
		fmt.Print("TIMEOUT: All sources were too slow")
	}
}

func player(name string, table chan string) {
	for {
		//wait for ball to come to you
		currentBallBearer := <-table
		fmt.Printf("%s hit the ball!\n", name)
		time.Sleep(300 * time.Millisecond)
		//send the ball back
		table <- currentBallBearer
	}
}

func fetchPrice(source string, c chan string) {
	//simulate random network latency
	latency := time.Duration(rand.Intn(900)+100) * time.Millisecond
	time.Sleep(latency)
	c <- fmt.Sprintf("Price from %s: 65000 (took %v)", source, latency)
}
