package concurrency

import (
    "fmt"
    "testing"
    "time"
)

func attackBuddy(target string) {
    time.Sleep(time.Second)
    fmt.Println("Throwing ninja stars at", target)
}

func TestGorouting(t *testing.T) {
    start := time.Now()

    defer func() {
        fmt.Println(time.Since(start))
    }()

    buddies := []string{"Ben", "Tom", "Jim", "Nio"}

    for _, buddy := range buddies {
        go attackBuddy(buddy)
    }

    time.Sleep(time.Second * 2)
}
