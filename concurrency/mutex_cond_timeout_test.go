package concurrency

import (
    "sync"
    "testing"
    "time"
)

func TestConditionTimeout(t *testing.T) {
    var (
        mu   sync.Mutex
        cond = sync.NewCond(&mu)
    )

    timer := time.NewTimer(5 * time.Second)
    defer timer.Stop()

    done := make(chan struct{})

    go func() {
        mu.Lock()
        defer mu.Unlock()
        cond.Wait()
        t.Log("I am awake")
        close(done)
    }()

    select {
    case <-done:
        t.Log("cond.Wait() was completed")
    case <-timer.C:
        t.Fatal("Test timed out, cond.Wait() was not completed")
    }
}

func TestConditionAwakeBeforeTimeout(t *testing.T) {
    var (
        mu   sync.Mutex
        cond = sync.NewCond(&mu)
    )

    timer := time.NewTimer(5 * time.Second)
    defer timer.Stop()

    done := make(chan struct{})

    go func() {
        mu.Lock()
        defer mu.Unlock()
        cond.Wait()
        t.Log("I am awake")
        close(done)
    }()

    time.Sleep(2 * time.Second)
    cond.Signal()

    select {
    case <-done:
        t.Log("cond.Wait() was completed")
    case <-timer.C:
        t.Fatal("Test timed out, cond.Wait() was not completed")
    }
}
