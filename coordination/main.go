package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}
var rwmutex = sync.RWMutex{}
var squares = map[int]int{}

func doSum(count int, val *int) {
	time.Sleep(time.Second)
	mutex.Lock()
	for i := 0; i < count; i++ {
		*val++
	}
	mutex.Unlock()
	waitGroup.Done()
}

func calculateSquares(max, iterations int) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Added value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value: %v = %v", val, squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}

func main() {
	// counter := 0
	rand.Seed(time.Now().UnixNano())

	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		// go doSum(5000, &counter)
		go calculateSquares(10, 5)
	}
	waitGroup.Wait()
	// Printfln("Total: %v", counter)
	Printfln("Cached values: %v", len(squares))
}
