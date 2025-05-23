package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func simulateTask(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
}

func main() {
	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	numTasks := 1_000_000
	var wg sync.WaitGroup
	wg.Add(numTasks)

	for range numTasks {
		go simulateTask(&wg)
	}

	wg.Wait()

	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)

	fmt.Printf("Total tasks: %d\n", numTasks)
	fmt.Printf("Memory used: %.2f MB\n", float64(memEnd.Alloc-memStart.Alloc)/1024/1024)
}
