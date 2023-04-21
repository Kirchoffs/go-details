package concurrency

import (
    "sync"
    "testing"
    "time"
)

func TestCondition(t *testing.T) {
    var (
        mu    sync.Mutex
        cond  = sync.NewCond(&mu)
        count int
    )

    increment := func() {
        mu.Lock()
        defer mu.Unlock()

        count++
        cond.Signal()
    }

    waitForCondition := func() {
        mu.Lock()
        defer mu.Unlock()

        for count < 10000 {
            cond.Wait()
            t.Log("I am awake and I see count is", count)
        }
    }

    done := make(chan struct{})
    go func() {
        defer close(done)
        waitForCondition()
    }()

    for i := 0; i < 10000; i++ {
        go increment()
    }

    select {
    case <-done:
        t.Log("Condition was met")
    case <-time.After(5 * time.Second):
        t.Fatal("Test timed out, condition was not met")
    }
}
