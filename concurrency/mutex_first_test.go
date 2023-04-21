package concurrency

import (
    "fmt"
    "sync"
    "testing"
)

var (
    lock   sync.Mutex
    beeper sync.WaitGroup
    count  int
)

func TestMutexFirstDemo(t *testing.T) {
    iterations := 1000
    beeper.Add(iterations)
    for i := 0; i < iterations; i++ {
        go increment()
    }

    beeper.Wait()
    fmt.Println("Result:", count)
}

func increment() {
    lock.Lock()
    count++
    beeper.Done()
    lock.Unlock()
}
