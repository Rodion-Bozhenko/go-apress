package main

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	countKey = iota
	sleepPeriodKey
)

var waitGroup = sync.WaitGroup{}
var mutex = sync.Mutex{}
var rwmutex = sync.RWMutex{}
var squares = map[int]int{}
var readyCond = sync.NewCond(rwmutex.RLocker())
var once = sync.Once{}

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

func generateSquares(max int) {
	// rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	// rwmutex.Unlock()
	// Printfln("Broadcasting condition")
	// readyCond.Broadcast()
	// waitGroup.Done()
}

func readSquares(id, max, iterations int) {
	once.Do(func() {
		generateSquares(max)
	})
	// readyCond.L.Lock()
	// for len(squares) == 0 {
	// 	readyCond.Wait()
	// }
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read Value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	// readyCond.L.Unlock()
	waitGroup.Done()
}

func processRequest(ctx context.Context, wg *sync.WaitGroup) {
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(sleepPeriod)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

func main() {
	// counter := 0
	// rand.Seed(time.Now().UnixNano())

	// numRoutines := 2
	// waitGroup.Add(numRoutines)
	// for i := 0; i < numRoutines; i++ {
	// go doSum(5000, &counter)
	// go calculateSquares(10, 5)
	// 	go readSquares(i, 10, 5)
	// }
	// waitGroup.Add(1)
	// go generateSquares(10)
	// waitGroup.Wait()
	// Printfln("Total: %v", counter)
	// Printfln("Cached values: %v", len(squares))

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	Printfln("Request dispatched...")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	ctx = context.WithValue(ctx, countKey, 4)
	ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond*250)
	go processRequest(ctx, &waitGroup)

	// time.Sleep(time.Second)
	// Printfln("Cancelling request")
	// cancel()

	waitGroup.Wait()
}
