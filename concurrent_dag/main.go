package main

import (
	"fmt"
	"sync"
)

type Graph struct {
	adj     map[int][]int
	visited map[int]bool
	mu      sync.Mutex
}

func NewGraph() *Graph {
	return &Graph{adj: make(map[int][]int), visited: make(map[int]bool)}
}

func (g *Graph) AddEdge(u, v int) {
	g.adj[u] = append(g.adj[u], v)
}

//Find cycles performs DFS

func (g *Graph) findCycles(u int, onStack map[int]bool, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	g.mu.Lock()
	if g.visited[u] {
		g.mu.Unlock()
		return
	}

	g.visited[u] = true
	g.mu.Unlock()
	onStack[u] = true

	for _, v := range g.adj[u] {
		if onStack[v] {
			results <- 1 /// cycle detected
		} else {
			//In real concurrent dfs, you'd be careful aboutt spawning too many
			//goroutines deep in the recursion. Here we call it synchronously
			//for path but could parellelize neighbor exploration
			g.mu.Lock()
			visitedV := g.visited[v]
			g.mu.Unlock()

			if !visitedV {
				//recurse
				newStack := make(map[int]bool)
				for k, val := range onStack {
					newStack[k] = val
				}
				g.dfsRecursive(v, newStack, results)
			}
		}
	}
	onStack[u] = false
}

// recurive dfs heloer
func (g *Graph) dfsRecursive(u int, onStack map[int]bool, results chan<- int) {
	g.mu.Lock()
	g.visited[u] = true
	g.mu.Unlock()
	onStack[u] = true

	for _, v := range g.adj[u] {
		if onStack[v] {
			results <- 1
		} else {
			g.mu.Lock()
			if !g.visited[v] {
				g.mu.Unlock()
				g.dfsRecursive(v, onStack, results)
			} else {
				g.mu.Unlock()
			}

		}
	}
	onStack[u] = false
}

func main() {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(3, 4)
	g.AddEdge(4, 3)
	// g.AddEdge(0,1

	results := make(chan int, 10)
	var wg sync.WaitGroup
	for node := range g.adj {
		wg.Add(1)
		go g.findCycles(node, make(map[int]bool), results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalCycles := 0
	for count := range results {
		totalCycles += count
	}
	fmt.Printf("Total Cycles found: %d\n", totalCycles)
}
