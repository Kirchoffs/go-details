package concurrency

import (
    "fmt"
    "testing"
    "time"
)

func attackFailed(target string, smokeSignal chan string) {
    time.Sleep(time.Second)
    fmt.Println("Throwing ninja stars at", target)
    smokeSignal <- "Mission failed"
}

func TestChannelSecondDemot(t *testing.T) {
    start := time.Now()

    defer func() {
        fmt.Println(time.Since(start))
    }()

    smokeSignal := make(chan string)
    buddy := "Ben"

    go attackFailed(buddy, smokeSignal)
    fmt.Println(<-smokeSignal)
}
