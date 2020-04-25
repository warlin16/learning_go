package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// PrintInfo about concurrency
func PrintInfo() {
	wg.Add(1)
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	go first()
	second()
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func first() {
	for i := 0; i < 10; i++ {
		fmt.Println("First is printing the index:", i)
	}
	wg.Done()
}

func second() {
	for i := 0; i < 10; i++ {
		fmt.Println("Second is printing the index:", i)
	}
}

// create a race condition on purpose

// RaceCondition creates a race condition on purpose
func RaceCondition() {
	start := time.Now()
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("goroutines", runtime.NumGoroutine())

	counter := 0

	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count:", counter)
	fmt.Println("Time took this long to execute", time.Since(start))
}

// A mutex is used to lock a goroutine for cases where you're sharing
// access to a variable and changing its value between goroutines so that
// you cant cause a race condition
