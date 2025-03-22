package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime package
	cpuNum := runtime.NumCPU()
	fmt.Println("Number of CPUs: ", cpuNum)
}
